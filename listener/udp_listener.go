package listener

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Start(buf_size int, port string, c chan []byte) {
	udpServer, err := net.ListenPacket("udp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 2048)
		_, addr, err := udpServer.ReadFrom(buf)
		fmt.Printf("Received from %s", addr)
		c <- buf
		if err != nil {
			continue
		}
		go response(udpServer, addr, buf)
	}

}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v", time)
	udpServer.WriteTo([]byte(responseStr), addr)
}
