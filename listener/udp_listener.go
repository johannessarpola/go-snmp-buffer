package listener

import (
	"net"

	g "github.com/gosnmp/gosnmp"
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/serdes"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type SnmpHandler struct {
	// TODO Add support for v3 at some point
	gosnmp *g.GoSNMP
	out    chan<- []byte // TODO add some serious tool for this
}

func (snmp *SnmpHandler) HandlePacket(pckt []byte) error {
	trap, err := snmp.gosnmp.UnmarshalTrap(pckt, false) // only supports v2 now
	p := models.NewPacket(trap)
	b, err := serdes.Encode(&p)
	if err != nil {
		logger.Info("Encoding failed!!")
		return err
	} else {
		// TODO Add some actual lib instead of chan
		snmp.out <- b // arr is copied to channel
	}
	return nil
}

func NewSnmpHandler(gosnmp *g.GoSNMP, out chan<- []byte) *SnmpHandler {
	h := &SnmpHandler{
		gosnmp: gosnmp,
		out:    out,
	}
	return h
}

func Start(port int, data *db.Data) {
	// TODO Cleanup
	defer ants.Release()
	ants, err := ants.NewPool(100)
	if err != nil {
		panic(err)
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf [4096]byte
	for {
		rlen, remote, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			logger.Warn("could not read packet", err)
		}
		logger.Infof("got trapdata from %s\n", remote.IP)
		pckt := buf[:rlen]

		err = ants.Submit(func() {
			rsc := make(chan []byte, 1)
			h := NewSnmpHandler(g.Default, rsc) // TODO quite ugly, refactor at some point
			h.HandlePacket(pckt)
			data.Append(<-rsc)
		})
		if err != nil {
			logger.Warn("Error in processing goroutine", err)
		}
	}

}
