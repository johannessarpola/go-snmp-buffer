package snmp

import (
	"net"

	"github.com/gosnmp/gosnmp"
	m "github.com/johannessarpola/go-network-buffer/pkg/metrics"
	"github.com/johannessarpola/go-network-buffer/pkg/models"
	s "github.com/johannessarpola/go-network-buffer/pkg/snmpdb"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// TODO Rename to be more clear and not snmp.ListenSnmp()
func ListenSnmp(port int, host string, data *s.SnmpDB) {
	// TODO Cleanup
	defer ants.Release()
	ants, err := ants.NewPool(100)
	if err != nil {
		panic(err)
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf [4096]byte
	dones := make(chan bool)
	defer close(dones)
	go m.MeasureRate(dones)
	for {
		rlen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			logger.Warn("could not read packet", err)
		}
		//logger.Debug("got trapdata from %s\n", remote.IP.String())
		pckt := buf[:rlen]

		// TODO Cleanup
		err = ants.Submit(func() {
			rsc := make(chan models.Element, 1)
			h := NewSnmpHandler(gosnmp.Default, rsc) // TODO quite ugly, refactor at some point
			err := h.HandlePacket(pckt)

			if err != nil {
				logger.Error("Could not handle packet", err)
			} else {
				err := data.Enqueue(<-rsc)
				if err != nil {
					logger.Error("Could not enqueue packet")
				}
			}
			dones <- true
		})
		if err != nil {
			logger.Warn("Error in processing goroutine", err)
		}
	}

}
