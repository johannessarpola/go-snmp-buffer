package listener

import (
	"fmt"
	"log"
	"net"
	"os"

	g "github.com/gosnmp/gosnmp"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/snmp"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Start(port string, out chan<- []byte) {
	tl := g.NewTrapListener()
	tl.OnNewTrap = func(s *g.SnmpPacket, u *net.UDPAddr) {

		logger.Infof("got trapdata from %s\n", u.IP)
		p := models.NewPacket(s)
		b, err := snmp.Encode(&p)

		if err != nil {
			logger.Info("Encoding failed!!")
		} else {
			out <- b // arr is copied to channel
		}

	}
	tl.Params = g.Default
	tl.Params.Logger = g.NewLogger(log.New(os.Stdout, "", 0)) // TODO use logrus
	err := tl.Listen(fmt.Sprintf("0.0.0.0:%s", port))

	if err != nil {
		logger.Fatal("Could not start listener")
	}

	defer tl.Close()

}
