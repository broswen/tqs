package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/broswen/tqs/internal/message"
	"github.com/broswen/tqs/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func PublishMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &PublishMessageRequest{}
		if err := render.Bind(r, request); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		message := &repository.Message{
			Topic:      request.Topic,
			Data:       request.Data,
			Attributes: request.Attributes,
		}
		err := service.Publish(message)
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}
		render.Render(w, r, &PublishMessageResponse{Id: message.Id})
	}
}

func ReceiveMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &ReceiveMessageRequest{}
		topic := chi.URLParam(r, "name")
		request.Topic = topic

		if request.Topic == "" {
			render.Render(w, r, ErrInvalidRequest(errors.New("topic name is missing")))
			return
		}
		messages, err := service.Receive(request.Topic)
		if err != nil {
			log.Println(err)
			render.Render(w, r, ErrInternalServer(err))
			return
		}
		render.Render(w, r, &ReceiveMessageResponse{
			Messages: messages,
		})
	}
}

func AckMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &AckMessageRequest{}
		id := chi.URLParam(r, "id")
		topic := chi.URLParam(r, "name")
		request.Id = id
		request.Topic = topic

		if request.Id == "" {
			render.Render(w, r, ErrInvalidRequest(errors.New("message id is missing")))
			return
		}

		if request.Topic == "" {
			render.Render(w, r, ErrInvalidRequest(errors.New("topic name is missing")))
			return
		}

		message := &repository.Message{
			Id:    request.Id,
			Topic: request.Topic,
		}
		err := service.Ack(message)
		if err != nil {
			log.Println(err)
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		render.Render(w, r, &AckMessageResponse{Id: message.Id})
	}
}
