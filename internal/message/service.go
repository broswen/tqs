package message

import (
	"time"

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
	if message.Expiration == 0 {
		message.Expiration = time.Now().Add(7 * 24 * time.Hour).Unix()
	}
	if message.Visible == 0 {
		message.Visible = time.Now().Unix()
	}
	err := ms.repo.SaveMessage(message)
	return err
}

func (ms MessageService) Receive(topic string, limit int, attributes map[string]string) ([]repository.Message, error) {
	messages, err := ms.repo.GetMessagesByTopic(topic, limit, attributes)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func (ms MessageService) Ack(message *repository.Message) error {
	message.Ack = time.Now().Unix()
	err := ms.repo.UpdateMessage(message)
	return err
}
