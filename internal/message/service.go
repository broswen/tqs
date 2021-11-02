package message

import (
	"github.com/broswen/tqs/internal/repository"
)

type MessageService struct {
	repo repository.MessageRepository
}

func New(repo repository.MessageRepository) (MessageService, error) {
	return MessageService{
		repo: repo,
	}, nil
}

func (ms MessageService) Publish(message *repository.Message) error {
	return nil
}

func (ms MessageService) Receive(topic string) ([]repository.Message, error) {
	return []repository.Message{}, nil
}

func (ms MessageService) Ack(id string) error {
	return nil
}
