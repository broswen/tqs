package repository

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMongoRepoInit(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	_, err := NewMongoMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}
}

func TestMongoRepoSave(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	repo, err := NewMongoMessageRepository()
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

	if m.Id.IsZero() {
		t.Fatalf("message Id is not set\n")
	}
	fmt.Println(m.Id)
}

func TestMongoRepoGet(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	repo, err := NewMongoMessageRepository()
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

	if m.Id.IsZero() {
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

func TestMongoRepoDelete(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	repo, err := NewMongoMessageRepository()
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

func TestMongoRepoReceive(t *testing.T) {
	os.Setenv("MONGODB_USER", "tqs")
	os.Setenv("MONGODB_PASS", "password")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "tqs")
	repo, err := NewMongoMessageRepository()
	if err != nil {
		t.Fatalf("init repo: %v\n", err)
	}

	topic := "test"

	err = repo.SaveMessage(&Message{Topic: topic, Data: "ok"})
	// should ignore non visible messages
	err = repo.SaveMessage(&Message{Topic: topic, Data: "non visible", Visible: time.Now().Add(time.Hour).Unix()})
	// should ignore ack'd messages
	err = repo.SaveMessage(&Message{Topic: topic, Data: "acked", Ack: time.Now().Unix()})
	// should ignore expired messages
	err = repo.SaveMessage(&Message{Topic: topic, Data: "expiration", Expiration: time.Now().Add(-1 * time.Second).Unix()})
	if err != nil {
		t.Fatalf("save message: %v\n", err)
	}

	messages, err := repo.GetMessagesByTopic(topic)
	if err != nil {
		t.Fatalf("get messages by topic: %v\n", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %d\n", len(messages))
	}

}
