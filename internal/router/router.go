package router

import "connect-meme/internal/subscriber"

type Router interface {
	listen()
	onConnection(subscriber subscriber.Subscriber)
}
