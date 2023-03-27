package main

import (
	"time"

	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/listener"
)

func main() {
	c := make(chan []byte)
	go listener.Start(1024, "9999", c)

	data := db.NewData("../../_tmp", "snmp_")
	go data.Connect(c)

	for {
		time.Sleep(1 * time.Second) // TODO Remove
	}
}
