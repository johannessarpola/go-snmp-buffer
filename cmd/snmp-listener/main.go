package main

import (
	"time"

	"github.com/johannessarpola/go-network-buffer/listener"
)

func main() {
	// Hello world, the web server

	go listener.Start(1024, "8081")
	for {
		time.Sleep(1 * time.Second) // TODO Remove
	}
}
