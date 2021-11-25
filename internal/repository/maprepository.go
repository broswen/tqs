package repository

import (
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapMessageRepository struct {
	topics map[string]map[string]Message
}

func NewMapMessageRepository() (MapMessageRepository, error) {
	return MapMessageRepository{
		topics: make(map[string]map[string]Message),
	}, nil
}

func (mr MapMessageRepository) SaveMessage(m *Message) error {
	_, ok := mr.topics[m.Topic]
	if !ok {
		mr.topics[m.Topic] = make(map[string]Message)
	}
	m.Expiration = time.Now().Add(7 * 24 * time.Hour).Unix()
	m.Visible = time.Now().Unix()
	topic := mr.topics[m.Topic]

	m.Id = primitive.NewObjectID()
	topic[m.Id.Hex()] = *m

	return nil
}

func (mr MapMessageRepository) UpdateMessage(m *Message) error {
	_, ok := mr.topics[m.Topic]
	if !ok {
		mr.topics[m.Topic] = make(map[string]Message)
	}
	topic := mr.topics[m.Topic]
	if _, ok = topic[m.Id.Hex()]; !ok {
		return errors.New("update message doesn't exist")
	}
	topic[m.Id.Hex()] = *m
	return nil
}

func (mr MapMessageRepository) GetMessage(m *Message) error {
	topic, ok := mr.topics[m.Topic]
	if !ok {
		return errors.New("no messages for topic found")
	}

	m2, ok := topic[m.Id.Hex()]
	if !ok {
		return errors.New("no messages with id found in topic")
	}
	*m = m2
	return nil
}

func (mr MapMessageRepository) DeleteMessage(m *Message) error {
	topic, ok := mr.topics[m.Topic]
	if !ok {
		return nil
	}

	delete(topic, m.Id.Hex())
	return nil
}

func (mr MapMessageRepository) GetMessagesByTopic(topicName string) ([]Message, error) {
	now := time.Now().Unix()
	topic, ok := mr.topics[topicName]
	if !ok {
		return []Message{}, nil
	}
	messages := make([]Message, 0)
	for _, v := range topic {
		// skip if acknowledged
		if v.Ack != 0 {
			continue
		}
		// skip if not visible
		if now < v.Visible {
			continue
		}
		// skip if expired
		if now >= v.Expiration {
			continue
		}

		v.Ack = time.Now().Unix()
		if err := mr.UpdateMessage(&v); err != nil {
			log.Printf("update message: %v\n", err)
		}

		messages = append(messages, v)
	}

	return messages, nil
}
