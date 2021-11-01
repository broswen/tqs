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

func (pm *PublishMessageRequest) Bind(r *http.Request) error {
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

func (rm *ReceiveMessageRequest) Bind(r *http.Request) error {
	return nil
}

type ReceiveMessageResponse struct {
	Messages []Message `json:"messages"`
}

func (rm ReceiveMessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Message struct {
	Id         string            `json:"id"`
	Topic      string            `json:"topic"`
	Attributes map[string]string `json:"attributes"`
	Data       string            `json:"data"`
}

type AckMessageRequest struct {
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
