package storage

import (
	"bytes"
	"connect-meme/internal/subscriber"
	"github.com/thoas/go-funk"
)

type InMemoryStorage struct {
	subscribers []subscriber.Subscriber
}
func NewInMemoryStorage() (storage *InMemoryStorage){
	return &InMemoryStorage{
		subscribers: []subscriber.Subscriber{},
	}
}

func (storage InMemoryStorage) RegisterSubscriber(subscriber subscriber.Subscriber) subscriber.Subscriber {
	storage.subscribers = append(storage.subscribers, subscriber)
	return subscriber
}

func (storage InMemoryStorage) findByClientId(clientId []byte) subscriber.Subscriber {
	return funk.Find(storage.subscribers, func(s subscriber.Subscriber) bool {
		return bytes.Equal(clientId, s.ClientId)
	}).(subscriber.Subscriber)
}

func (storage InMemoryStorage) findByUserId(userId string) subscriber.Subscriber {
	return funk.Find(storage.subscribers, func(s subscriber.Subscriber) bool {
		return userId == s.UserId
	}).(subscriber.Subscriber)
}
