package listener

import (
	"net"

	g "github.com/gosnmp/gosnmp"
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/serdes"
	u "github.com/johannessarpola/go-network-buffer/utils"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type SnmpHandler struct {
	// TODO Add support for v3 at some point
	gosnmp *g.GoSNMP
	out    chan<- models.Element // TODO add some serious tool for this
}

func (snmp *SnmpHandler) HandlePacket(pckt []byte) error {
	trap, err := snmp.gosnmp.UnmarshalTrap(pckt, false) // only supports v2 now
	if err != nil {
		panic(err) // TODO
	}
	p := models.NewPacket(trap)
	b, err := serdes.Encode(&p)
	if err != nil {
		logger.Info("Encoding failed!!")
		return err
	} else {
		// TODO Add some actual lib instead of chan
		snmp.out <- models.NewElement(b) // arr is copied to channel
	}
	return nil
}

func NewSnmpHandler(gosnmp *g.GoSNMP, out chan<- models.Element) *SnmpHandler {
	h := &SnmpHandler{
		gosnmp: gosnmp,
		out:    out,
	}
	return h
}

func Start(port int, data *db.Database) {
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
	dones := make(chan bool)
	defer close(dones)
	go u.MeasureRate(dones)
	for {
		rlen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			logger.Warn("could not read packet", err)
		}
		//logger.Debug("got trapdata from %s\n", remote.IP.String())
		pckt := buf[:rlen]

		err = ants.Submit(func() {
			rsc := make(chan models.Element, 1)
			h := NewSnmpHandler(g.Default, rsc) // TODO quite ugly, refactor at some point
			h.HandlePacket(pckt)
			data.RingDB.Enqueue(<-rsc)
			dones <- true
		})
		if err != nil {
			logger.Warn("Error in processing goroutine", err)
		}
	}

}
