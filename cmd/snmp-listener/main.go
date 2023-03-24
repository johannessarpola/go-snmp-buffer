package main

import (
	"time"

	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/listener"
)

func main() {
	c := make(chan []byte)
	go listener.Start(1024, "8081", c)

	data := db.Data{}
	go data.Start("../../_tmp", c)

	for {
		time.Sleep(1 * time.Second) // TODO Remove
	}
}
