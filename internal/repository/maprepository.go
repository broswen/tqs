package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	topic := mr.topics[m.Topic]

	if m.Id == "" {
		m.Id = fmt.Sprintf("%d", rand.Intn(1000))
	}
	topic[m.Id] = *m

	return nil
}

func (mr MapMessageRepository) GetMessage(m *Message) error {
	topic, ok := mr.topics[m.Topic]
	if !ok {
		return errors.New("no messages for topic found")
	}

	m2, ok := topic[m.Id]
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

	delete(topic, m.Id)
	return nil
}

func (mr MapMessageRepository) GetMessagesByTopic(topicName string) ([]Message, error) {
	topic, ok := mr.topics[topicName]
	if !ok {
		return []Message{}, nil
	}
	messages := make([]Message, 0)
	for _, v := range topic {
		// skip if acknowledged
		if (v.Ack != time.Time{}) {
			continue
		}
		// skip if not visibile
		if time.Now().Before(v.Visible) {
			continue
		}
		// skip if expired
		if (v.Expiration != time.Time{}) && time.Now().After(v.Expiration) {
			continue
		}
		messages = append(messages, v)
	}
	return messages, nil
}
