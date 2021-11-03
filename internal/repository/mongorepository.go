package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoMessageRepository struct {
	mongoClient *mongo.Client
	messages    *mongo.Collection
}

func NewMongoMessageRepository() (MongoMessageRepository, error) {
	user := os.Getenv("MONGODB_USER")
	pass := os.Getenv("MONGODB_PASS")
	host := os.Getenv("MONGODB_HOST")
	port := os.Getenv("MONGODB_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return MongoMessageRepository{}, err
	}
	err = client.Ping(context.TODO(), &readpref.ReadPref{})
	if err != nil {
		return MongoMessageRepository{}, err
	}

	dbName := os.Getenv("MONGODB_DB")
	messageCollection := client.Database(dbName).Collection("messages")

	return MongoMessageRepository{
		mongoClient: client,
		messages:    messageCollection,
	}, nil
}

func (mr MongoMessageRepository) SaveMessage(message *Message) error {
	message.Expiration = time.Now().Add(7 * 24 * time.Hour)
	message.Visible = time.Now()
	result, err := mr.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	message.Id = fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (mr MongoMessageRepository) GetMessage(message *Message) error {
	oid, err := primitive.ObjectIDFromHex(message.Id)
	if err != nil {
		return err
	}
	result := mr.messages.FindOne(context.TODO(), bson.D{{"topic", message.Topic}, {"_id", oid}})
	if result.Err() != nil {
		return result.Err()
	}
	var m2 Message
	err = result.Decode(&m2)

	if err != nil {
		return err
	}

	*message = m2

	return nil
}

func (mr MongoMessageRepository) DeleteMessage(message *Message) error {
	oid, err := primitive.ObjectIDFromHex(message.Id)
	if err != nil {
		return err
	}
	_, err = mr.messages.DeleteOne(context.TODO(), bson.D{{"topic", message.Topic}, {"_id", oid}})
	if err != nil {
		return err
	}
	return nil
}

func (mr MongoMessageRepository) GetMessagesByTopic(topic string) ([]Message, error) {
	cursor, err := mr.messages.Find(context.TODO(), bson.D{{"topic", topic}})
	if err != nil {
		return []Message{}, err
	}

	messages := make([]Message, 0)

	for cursor.Next(context.TODO()) {
		var m2 Message
		if err = cursor.Decode(&m2); err != nil {
			continue
		}

		now := time.Now()

		// TODO mongodb bson does not handle time.Time properly
		if (m2.Ack != time.Time{}) {
			continue
		}
		// skip if not visibile
		if now.Before(m2.Visible) {
			continue
		}
		// skip if expired
		if (m2.Expiration != time.Time{}) && now.After(m2.Expiration) {
			continue
		}
		messages = append(messages, m2)
	}

	return messages, nil
}
