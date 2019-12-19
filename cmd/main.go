package main

import (
	"connect-meme/pkg/router"
)
func main() {
	router.NewPublisher("tcp://*:5555")
}
