package listener

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"

	g "github.com/gosnmp/gosnmp"
	"github.com/johannessarpola/go-network-buffer/models"
)

func Start(buf_size int, port string, c chan []byte) {
	var buf bytes.Buffer // Stand-in for a network connection
	enc := gob.NewEncoder(&buf)

	tl := g.NewTrapListener()
	tl.OnNewTrap = func(s *g.SnmpPacket, u *net.UDPAddr) {

		log.Printf("got trapdata from %s\n", u.IP)
		err := enc.Encode(models.NewPacket(s))
		if err != nil {
			println("Encoding failed!!")
		}
		c <- buf.Bytes() // arr is copied to channel

	}
	tl.Params = g.Default
	tl.Params.Logger = g.NewLogger(log.New(os.Stdout, "", 0))
	err := tl.Listen(fmt.Sprintf("0.0.0.0:%s", port))

	if err != nil {
		log.Fatal("Could not start listener")
	}

	defer tl.Close()

}
