# Tiny Queue Service (TQS)

Tiny Queue Service is a simple queue service for asynchronous message communication between systems.

Created with Go and MongoDB.

## Features
- Publishing messages to specific topics
- Receiving messages from specific topics
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

Send a `GET` request to the `/topic/{name}` endpoint to query for messages in a specific `topic`.

Use the `limit` property to specify how many to retrieve, and filter messages by the `attributes` property.
```json
{
  "topic": "topic name",
  "limit": 4
  "attributes": {
    "key": "value"
  }
}
```

Send a `PUT` request to the `/topic/{name}/{id}` endpoint to acknowledge a message and remove it from the queue.
```json
{
  "topic": "topic name",
  "id": "message id"
}
```

## Todo
- [ ] handle mongodb and time.Time type conversion using int64
- [ ] create helm chart for k8s deploy
- [ ] handle filtering by attributes, only 1 layer simple [string]string map allowed.