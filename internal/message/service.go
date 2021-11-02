package message

import "time"

type Message struct {
	Id         string            `json:"id"`
	Topic      string            `json:"topic"`
	Attributes map[string]string `json:"attributes"`
	Data       string            `json:"data"`
	Visibility time.Time         `json:"-"` // when the message becomes visible
	Ack        time.Time         `json:"-"` // when the message was acknowledged
	Expiration time.Time         `json:"-"` // when the message will expire
}

type MessageService struct {
}

func New() (MessageService, error) {
	return MessageService{}, nil
}

func (ms MessageService) Publish(message *Message) error {
	return nil
}

func (ms MessageService) Receive(topic string) ([]Message, error) {
	return []Message{}, nil
}

func (ms MessageService) Ack(id string) error {
	return nil
}
