# Tiny Queue Service (TQS)

Tiny Queue Service is a simple queue service for asynchronous message communication between systems.

Created with Go and MongoDB.

## Features
- Publishing messages to specific topics
- Receiving messages from specific topics
- Explicit acknowledgement of messages
- Filtering messages on message attributes
- Acknowledgement and visibility timeouts

## Usage
Send a `POST` request to the `/message` endpoint to publish a message to a specific `topic`.
```json
{
  "topic": "topic name",
  "attributes": {
    "key": "value"
  },
  "data": "message data"
}
```

Send a `POST` request to the `/topic` endpoint to query for messages in a specific `topic`.

Use the `limit` property to specify how many messages to retrieve, and filter messages by the `attributes` property.
```json
{
  "topic": "topic name",
  "limit": 4
  "attributes": {
    "key": "value"
  }
}
```

Send a `PUT` request to the `/message/{id}` endpoint to acknowledge a message and remove it from the queue.

## MongoDB Data Model
```json
{
  "_id": "508f191e810c19729de860ea",  // objectid generated by mongo
  "topic": "topic1",                  // name of the topic
  "attributes": {                     // simple [string]string map of attributes
    "key": "value"
  },
  "data": "message",                  // message data
  "visible": 1637462159,              // when the message will become visible
  "ack": 1637462159,                  // when the message was acknowledged
  "expiration": 1637462159            // when the message will expire
}
```

## Todo
- [x] handle mongodb and time.Time type conversion using int64
- [x] create k8s deployment/files for simple hosting
- [x] refactor endpoints to be more succinct
- [x] implement limit in repo
- [x] create helm chart for automated k8s deploy
- [x] handle filtering by attributes
- [ ] refactor and add more tests
