package repository

import (
	"fmt"
	"testing"
	"time"
)

func TestMapRepInit(t *testing.T) {
	_, err := NewMapMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}
}

func TestMapRepoSave(t *testing.T) {
	repo, err := NewMapMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}
	m := &Message{
		Topic: "test",
		Data:  "test",
	}
	err = repo.SaveMessage(m)
	if err != nil {
		t.Fatalf("save mesage: %v\n", err)
	}

	if m.Id == "" {
		t.Fatalf("message Id is not set\n")
	}
}

func TestMapRepoGet(t *testing.T) {
	repo, err := NewMapMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}

	data := "test"
	m := &Message{
		Topic: "test",
		Data:  data,
	}
	err = repo.SaveMessage(m)
	if err != nil {
		t.Fatalf("save message: %v\n", err)
	}

	if m.Id == "" {
		t.Fatalf("message Id is not set\n")
	}

	m2 := &Message{
		Topic: m.Topic,
		Id:    m.Id,
	}

	err = repo.GetMessage(m2)
	if err != nil {
		t.Fatalf("get message: %v\n", err)
	}

	if m2.Data != data {
		fmt.Printf("%#v\n", m2)
		t.Fatalf("message data was %s, but wanted %s\n", m2.Data, data)
	}

}

func TestMapRepoDelete(t *testing.T) {
	repo, err := NewMapMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}
	m := &Message{
		Topic: "test",
		Data:  "test",
	}
	err = repo.SaveMessage(m)
	if err != nil {
		t.Fatalf("save message: %v\n", err)
	}

	err = repo.DeleteMessage(m)
	if err != nil {
		t.Fatalf("delete mesage: %v\n", err)
	}
}

func TestMapRepoReceive(t *testing.T) {
	repo, err := NewMapMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}

	data := "test"
	topic := "test"

	repo.SaveMessage(&Message{Topic: topic, Data: data})
	repo.SaveMessage(&Message{Topic: "wrong", Data: data})
	// should ignore ack'd messages
	repo.SaveMessage(&Message{Topic: "wrong", Data: data, Ack: time.Now()})
	// should ignore expired messages
	repo.SaveMessage(&Message{Topic: "wrong", Data: data, Expiration: time.Now().Add(-1 * time.Second)})

	messages, err := repo.GetMessagesByTopic(topic)
	if err != nil {
		t.Fatalf("get messages by topic: %v\n", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %d\n", len(messages))
	}

}
