package repository

import (
	"context"
	"fmt"
	"log"
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
	
	result, err := mr.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (mr MongoMessageRepository) UpdateMessage(message *Message) error {
	update := bson.D{}
	if message.Visible != 0 {
		update = append(update, bson.E{"$set", bson.D{{"visible", message.Visible}}})
	}
	if message.Expiration != 0 {
		update = append(update, bson.E{"$set", bson.D{{"expiration", message.Expiration}}})
	}
	if message.Ack != 0 {
		update = append(update, bson.E{"$set", bson.D{{"ack", message.Ack}}})
	}
	_, err := mr.messages.UpdateByID(context.TODO(), message.Id, update)
	if err != nil {
		return err
	}
	return nil
}

func (mr MongoMessageRepository) GetMessage(message *Message) error {
	result := mr.messages.FindOne(context.TODO(), bson.D{{"topic", message.Topic}, {"_id", message.Id}})
	if result.Err() != nil {
		return result.Err()
	}
	var m2 Message
	err := result.Decode(&m2)

	if err != nil {
		return err
	}

	*message = m2

	return nil
}

func (mr MongoMessageRepository) DeleteMessage(message *Message) error {
	_, err := mr.messages.DeleteOne(context.TODO(), bson.D{{"topic", message.Topic}, {"_id", message.Id}})
	if err != nil {
		return err
	}
	return nil
}

func (mr MongoMessageRepository) GetMessagesByTopic(topic string, limit int, attributes map[string]string) ([]Message, error) {
	now := time.Now().Unix()
	// only get messages where not ack'd
	// visible is less than now
	// expiration is greater than now
	filter := bson.D{
		{"topic", topic},
		{"ack", 0},
		{"visible", bson.D{{"$lte", now}}},
		{"expiration", bson.D{{"$gt", now}}},
	}
	// convert attributes map to mongo query filters
	filters := MapToFilter(attributes)
	// append converted filters
	filter = append(filter, filters...)
	cursor, err := mr.messages.Find(context.TODO(), filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return []Message{}, err
	}

	messages := make([]Message, 0)

	for cursor.Next(context.TODO()) {
		// update each message with a new visible time
		var m2 Message
		if err = cursor.Decode(&m2); err != nil {
			log.Printf("error decoding: %v\n", err)
			continue
		}
		m2.Visible = time.Now().Add(1 * time.Minute).Unix()
		if err = mr.UpdateMessage(&m2); err != nil {
			log.Printf("error updating visibility timeout: %v\n", err)
			continue
		}
		messages = append(messages, m2)
	}

	return messages, nil
}

func MapToFilter(attributes map[string]string) []bson.E {
	filters := make([]bson.E, 0)
	for k, v := range attributes {
		e := bson.E{fmt.Sprintf("attributes.%s", k), v}
		filters = append(filters, e)
	}
	return filters
}