package handlers

import "net/http"

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

type PublishMessageResponse struct {
	Id string `json:"id"`
}

type ReceiveMessageRequest struct {
	Topic      string            `json:"topic"`
	Limit      int               `json:"limit"`
	Attributes map[string]string `json:"attributes"`
}

type ReceiveMessageResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Id         string            `json:"id"`
	Topic      string            `json:"topic"`
	Attributes map[string]string `json:"attributes"`
	Data       string            `json:"data"`
}

type AckMessageRequest struct {
}

type AckMessageResponse struct {
	Id string `json:"id"`
}
