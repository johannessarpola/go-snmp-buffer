package listener

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func Start(buf_size int, port string, c chan []byte) {
	udpServer, err := net.ListenPacket("udp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		fmt.Printf("Received %s from %s", strings.TrimSpace(string(buf)), addr)
		c <- buf
		if err != nil {
			continue
		}
		go response(udpServer, addr, buf)
	}

}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))
	udpServer.WriteTo([]byte(responseStr), addr)
}
