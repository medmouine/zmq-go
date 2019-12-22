package router

import (
	. "connect-meme/internal/storage"
	. "connect-meme/internal/subscriber"
	"connect-meme/pkg/util"
	"encoding/json"
	zmq "gopkg.in/zeromq/goczmq.v4"
	"log"
)

type ZMQRouter struct {
	channeler   *zmq.Channeler
	storage Storage
}

func NewZMQRouter(host string, storage Storage, defaultNodes []string) (router *ZMQRouter){
	channeler := zmq.NewRouterChanneler(host)

	return &ZMQRouter{channeler:channeler, storage:storage}
}

func (router *ZMQRouter) Destroy() {
	router.channeler.Destroy()
}

func (router *ZMQRouter) listen() {
	defer router.Destroy()
	for {
		event := <-router.channeler.RecvChan
		clientId := event[0]
		req := router.parseRequest(event)
		switch req.RequestType {
		case Connection:
			router.onConnection(Subscriber{
				ClientId: clientId,
				UserId:   req.Id,
			})
		}
	}
}

func (router *ZMQRouter) onConnection(sub Subscriber) {
	router.storage.RegisterSubscriber(sub)
	response := util.ToJson(Message{
		Code:    Accepted,
		Message: "Connection to publisher successful",
	})

	router.channeler.SendChan <- [][]byte{sub.ClientId, response.([]byte)}
	log.Println("New connection to publisher id : " + sub.UserId + " client id : " + string(sub.ClientId))
}


func (router *ZMQRouter) parseRequest(event [][]byte) Request {
	var req Request
	err := json.Unmarshal(event[1], &req)
	if err != nil {
		log.Fatal("Could not decode json : " + err.Error())
	}
	return req
}
