package main

import (
	"connect-meme/internal/router"
	"connect-meme/internal/storage"
	"gopkg.in/zeromq/goczmq.v4"
	"log"
)
func main() {
	imStorage := storage.NewInMemoryStorage()
	router.NewZMQRouter("tcp://*:5555", imStorage, []string{})


	// Create a dealer channeler and connect it to the routerImpl.
	dealer := goczmq.NewDealerChanneler("tcp://127.0.0.1:5555")
	defer dealer.Destroy()
	log.Println("dealer created and connected")

	dealer.SendChan <- [][]byte{[]byte(
		"{\"Id\": \"123\", \"RequestType\": \"CONNECT\", \"Message\": \"\"}")}

	log.Println("dealer sent '123'")
	for {
		reply := <- dealer.RecvChan
		log.Printf("dealer received '%s'", string(reply[0]))
	}
}
