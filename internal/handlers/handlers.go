package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func PublishMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &PublishMessageResponse{})
	}
}

func ReceiveMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &ReceiveMessageResponse{})
	}
}

func AckMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &AckMessageResponse{})
	}
}
