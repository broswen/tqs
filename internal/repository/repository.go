package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Topic      string             `json:"topic" bson:"topic"`
	Attributes map[string]string  `json:"attributes" bson:"attributes"`
	Data       string             `json:"data" bson:"data"`
	Visible    int64              `json:"-" bson:"visible"`    // when the message becomes visible
	Ack        int64              `json:"-" bson:"ack"`        // when the message was acknowledged
	Expiration int64              `json:"-" bson:"expiration"` // when the message will expire
}

type MessageRepository interface {
	SaveMessage(message *Message) error
	GetMessage(message *Message) error
	DeleteMessage(message *Message) error
	GetMessagesByTopic(topic string, limit int, attributes map[string]string) ([]Message, error)
	UpdateMessage(message *Message) error
}
