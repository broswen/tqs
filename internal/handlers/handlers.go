package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func PublishMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &MessageResponse{Message: "OK"})
	}
}

func ReceiveMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &MessageResponse{Message: "OK"})
	}
}

func AckMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &MessageResponse{Message: "OK"})
	}
}
