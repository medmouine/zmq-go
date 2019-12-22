package storage

import (
	"connect-meme/internal/subscriber"
)

type Storage interface {
	RegisterSubscriber(subscriber subscriber.Subscriber) subscriber.Subscriber
	findByClientId(clientId []byte) subscriber.Subscriber
	findByUserId(userId string) subscriber.Subscriber
}
