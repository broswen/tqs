package message

import (
	"log"
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
	if (message.Expiration == time.Time{}) {
		// default expiration of 1 week
		message.Expiration = time.Now().Add(24 * 7 * time.Hour)
	}
	err := ms.repo.SaveMessage(message)
	return err
}

func (ms MessageService) Receive(topic string, limit int) ([]repository.Message, error) {
	messages, err := ms.repo.GetMessagesByTopic(topic)
	m2 := make([]repository.Message, 0)
	count := 0
	for _, m := range messages {
		if count >= limit {
			break
		}
		// set visibility timeout to 10 minutes
		m.Visible = time.Now().Add(10 * time.Minute)
		err := ms.repo.SaveMessage(&m)
		if err != nil {
			log.Printf("update visibility: %v\n", err)
		}
		m2 = append(m2, m)
		count++
	}
	return m2, err
}

func (ms MessageService) Ack(message *repository.Message) error {
	err := ms.repo.GetMessage(message)
	if err != nil {
		return err
	}
	message.Ack = time.Now()

	err = ms.repo.SaveMessage(message)
	return err
}
