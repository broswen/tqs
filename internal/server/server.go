package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/broswen/tqs/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

type Server interface {
	Start() error
}

type ChiServer struct {
	logger zerolog.Logger
	router chi.Router
}

func New() (ChiServer, error) {

	logger := httplog.NewLogger("tqs", httplog.Options{
		JSON: true,
	})
	server := ChiServer{
		logger: logger,
		router: chi.NewRouter(),
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

	// health check
	s.router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	s.router.Post("/message", handlers.PublishMessageHandler())
	s.router.Get("/topic/{name}", handlers.ReceiveMessageHandler())
	s.router.Put("/message/{id}", handlers.AckMessageHandler())
}
