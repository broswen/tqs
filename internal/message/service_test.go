package message

import (
	"fmt"
	"os"
	"testing"

	"github.com/broswen/tqs/internal/repository"
)

func TestPublish(t *testing.T) {
	mmr, err := repository.NewMapMessageRepository()
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}
	service, err := New(mmr)
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}

	topic := "test"
	data := "test"
	message := &repository.Message{
		Topic: topic,
		Data:  data,
	}

	err = service.Publish(message)
	if err != nil {
		t.Fatalf("Publish: %v\n", err)
	}
}

func TestReceive(t *testing.T) {
	mmr, err := repository.NewMapMessageRepository()
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}
	service, err := New(mmr)
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}

	topic := "test"
	data := "test"

	service.Publish(&repository.Message{Topic: topic, Data: data})
	service.Publish(&repository.Message{Topic: "a", Data: data})
	service.Publish(&repository.Message{Topic: "b", Data: data})

	messages, err := service.Receive(topic, 10)
	if err != nil {
		t.Fatalf("Receive: %v\n", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %d\n", len(messages))
	}

	messages, err = service.Receive(topic, 10)
	if err != nil {
		t.Fatalf("Receive: %v\n", err)
	}

	if len(messages) != 0 {
		t.Fatalf("expected 0 messages but got %d\n", len(messages))
	}
}

func TestAck(t *testing.T) {
	mmr, err := repository.NewMapMessageRepository()
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}
	service, err := New(mmr)
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}

	topic := "test"
	data := "test"

	service.Publish(&repository.Message{Topic: topic, Data: data})
	service.Publish(&repository.Message{Topic: "a", Data: data})
	service.Publish(&repository.Message{Topic: "b", Data: data})

	messages, err := service.Receive(topic, 10)
	if err != nil {
		t.Fatalf("Receive: %v\n", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %d\n", len(messages))
	}

	for _, m := range messages {
		err = service.Ack(&m)
		fmt.Printf("%#v\n", m)
		if err != nil {
			t.Fatalf("Ack: %v\n", err)
		}
	}
}

func TestReceiveMongo(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	mmr, err := repository.NewMongoMessageRepository()
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}
	service, err := New(mmr)
	if err != nil {
		t.Fatalf("NewMapMessageRepository: %v\n", err)
	}

	topic := "test"
	data := "test"

	service.Publish(&repository.Message{Topic: topic, Data: data})
	service.Publish(&repository.Message{Topic: "a", Data: data})
	service.Publish(&repository.Message{Topic: "b", Data: data})

	messages, err := service.Receive(topic, 10)
	if err != nil {
		t.Fatalf("Receive: %v\n", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %d\n", len(messages))
	}

	messages, err = service.Receive(topic, 10)
	if err != nil {
		t.Fatalf("Receive: %v\n", err)
	}

	if len(messages) != 0 {
		t.Fatalf("expected 0 messages but got %d\n", len(messages))
	}
}
