package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/broswen/tqs/internal/message"
	messages "github.com/broswen/tqs/internal/message"
)

type MessageRepository interface {
	SaveMessage(message *messages.Message) error
	GetMessage(message *messages.Message) error
	DeleteMessage(message *messages.Message) error
	GetMessagesByTopic(topic string) ([]messages.Message, error)
}

type MongoMessageRepository struct {
}

func NewMongoMessageRepository() (MongoMessageRepository, error) {
	return MongoMessageRepository{}, nil
}

func (m MongoMessageRepository) SaveMessage(message *messages.Message) error {
	return nil
}

func (m MongoMessageRepository) Getmessage(message *messages.Message) error {
	return nil
}

func (m MongoMessageRepository) DeleteMessage(message *messages.Message) error {
	return nil
}

func (m MongoMessageRepository) GetMessagesByTopic(topic string) ([]messages.Message, error) {
	return []messages.Message{}, nil
}

type MapMessageRepository struct {
	topics map[string]map[string]messages.Message
}

func NewMapMessageRepository() (MapMessageRepository, error) {
	return MapMessageRepository{
		topics: make(map[string]map[string]messages.Message),
	}, nil
}

func (m MapMessageRepository) SaveMessage(message *messages.Message) error {
	_, ok := m.topics[message.Topic]
	if !ok {
		m.topics[message.Topic] = make(map[string]messages.Message)
	}
	topic := m.topics[message.Topic]
	message.Id = fmt.Sprintf("%d", rand.Intn(1000))
	topic[message.Id] = *message

	return nil
}

func (m MapMessageRepository) GetMessage(message *message.Message) error {
	topic, ok := m.topics[message.Topic]
	if !ok {
		return errors.New("no messages for topic found")
	}

	m2, ok := topic[message.Id]
	if !ok {
		return errors.New("no messages with id found in topic")
	}
	*message = m2
	return nil
}

func (m MapMessageRepository) DeleteMessage(message *message.Message) error {
	topic, ok := m.topics[message.Topic]
	if !ok {
		return nil
	}

	delete(topic, message.Id)
	return nil
}

func (m MapMessageRepository) GetMessagesByTopic(topicName string) ([]message.Message, error) {
	topic, ok := m.topics[topicName]
	if !ok {
		return []message.Message{}, nil
	}
	messages := make([]messages.Message, 0)
	for _, v := range topic {
		if (v.Ack != time.Time{}) {

			continue
		}

		if (v.Expiration != time.Time{}) && time.Now().After(v.Expiration) {
			continue
		}
		messages = append(messages, v)
	}
	return messages, nil
}
