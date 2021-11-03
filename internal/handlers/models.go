package handlers

import (
	"errors"
	"net/http"

	"github.com/broswen/tqs/internal/repository"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func (mr MessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PublishMessageRequest struct {
	Topic      string            `json:"topic"`
	Attributes map[string]string `json:"attributes"`
	Data       string            `json:"data"`
}

func (pm *PublishMessageRequest) Bind(r *http.Request) error {
	if pm.Topic == "" {
		return errors.New("topic name is missing")
	}
	return nil
}

type PublishMessageResponse struct {
	Id string `json:"id"`
}

func (pm PublishMessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ReceiveMessageRequest struct {
	Topic      string            `json:"topic"`
	Limit      int               `json:"limit"`
	Attributes map[string]string `json:"attributes"`
}

type ReceiveMessageResponse struct {
	Messages []repository.Message `json:"messages"`
}

func (rm ReceiveMessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AckMessageRequest struct {
	Topic string `json:"topic"`
	Id    string `json:"id"`
}

func (am *AckMessageRequest) Bind(r *http.Request) error {
	return nil
}

type AckMessageResponse struct {
	Id string `json:"id"`
}

func (am AckMessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
