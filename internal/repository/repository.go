package repository

import (
	"time"
)

type Message struct {
	Id         string            `json:"id" bson:"_id,omitempty"`
	Topic      string            `json:"topic" bson:"topic"`
	Attributes map[string]string `json:"attributes" bson:"attributes"`
	Data       string            `json:"data" bson:"data"`
	Visible    time.Time         `json:"-" bson:"visible"`    // when the message becomes visible
	Ack        time.Time         `json:"-" bson:"ack"`        // when the message was acknowledged
	Expiration time.Time         `json:"-" bson:"expiration"` // when the message will expire
}

type MessageRepository interface {
	SaveMessage(message *Message) error
	GetMessage(message *Message) error
	DeleteMessage(message *Message) error
	GetMessagesByTopic(topic string) ([]Message, error)
}
