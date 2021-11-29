package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/broswen/tqs/internal/handlers"
	"github.com/broswen/tqs/internal/message"
	"github.com/broswen/tqs/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

type Server interface {
	Start() error
}

type ChiServer struct {
	logger         zerolog.Logger
	router         chi.Router
	messageService message.MessageService
}

func New() (ChiServer, error) {

	repo, err := repository.NewMongoMessageRepository()
	if err != nil {
		return ChiServer{}, err
	}
	service, err := message.New(repo)
	if err != nil {
		return ChiServer{}, err
	}

	logger := httplog.NewLogger("tqs", httplog.Options{
		JSON: true,
	})
	server := ChiServer{
		logger:         logger,
		router:         chi.NewRouter(),
		messageService: service,
	}
	server.SetRoutes()
	return server, nil
}

func (s ChiServer) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.logger.Info().Msgf("Starting chi server on :%s ...", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}

func (s ChiServer) SetRoutes() {
	s.router.Use(httplog.RequestLogger(s.logger))
	s.router.Use(render.SetContentType(render.ContentTypeJSON))
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	// health check
	s.router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	s.router.Post("/message", handlers.PublishMessageHandler(s.messageService))
	s.router.Put("/message/{id}", handlers.AckMessageHandler(s.messageService))
	s.router.Post("/topic", handlers.ReceiveMessageHandler(s.messageService))
}
