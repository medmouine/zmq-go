package router

import (
	"connect-meme/pkg/util"
	zmq "gopkg.in/zeromq/goczmq.v4"
	"log"
)


const (
	Accepted int = 202
)

const (
	Connection string = "CONNECT"
)

type message struct {
	Code int
	Message string
}

type request struct {
	id string
	requestType string
	message string
}

type subscriber struct {
	clientId []byte
	userId string
}
type ZMQPublisher struct {
	channeler   *zmq.Channeler
	Subscribers []subscriber
}

func NewPublisher(host string) (publisher *ZMQPublisher) {
	channeler := zmq.NewRouterChanneler(host)

	publisher = &ZMQPublisher{
		channeler:   channeler,
		Subscribers: []subscriber{},
	}

	go publisher.listen()

	return
}

func (publisher *ZMQPublisher) Destroy() {
	publisher.channeler.Destroy()
	publisher.Subscribers = []subscriber{}
}

func (publisher *ZMQPublisher) listen() {
	defer publisher.Destroy()
	for {
		event := <-publisher.channeler.RecvChan
		clientId := event[0]
		request := util.FromJson(event[1]).(request)
		switch request.requestType {
		case Connection:
			publisher.onConnection(subscriber{
				clientId: clientId,
				userId:   request.id,
			})
		}
	}
}

func (publisher *ZMQPublisher) onConnection(subscriber subscriber) {
	publisher.Subscribers = append(publisher.Subscribers, subscriber)
	response := util.ToJson(message{
		Code:    Accepted,
		Message: "Connection to publisher successful",
	})
	publisher.channeler.SendChan <- [][]byte{subscriber.clientId, response.([]byte)}
	log.Println("New connection to publisher")
	println(publisher.Subscribers)
}
