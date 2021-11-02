package handlers

import (
	"net/http"

	"github.com/broswen/tqs/internal/message"
	"github.com/go-chi/render"
)

func PublishMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &PublishMessageResponse{})
	}
}

func ReceiveMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &ReceiveMessageResponse{})
	}
}

func AckMessageHandler(service message.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &AckMessageResponse{})
	}
}
