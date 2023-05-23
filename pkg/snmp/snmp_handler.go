package snmp

import (
	g "github.com/gosnmp/gosnmp"
	_ "github.com/johannessarpola/go-network-buffer/pkg/logging"
	"github.com/johannessarpola/go-network-buffer/pkg/models"
	"github.com/johannessarpola/go-network-buffer/pkg/serdes"
	"github.com/sirupsen/logrus"
)

type Mapper interface {
	Handle(pckt []byte) error
}

type SnmpHandler struct {
	// TODO Add support for v3 at some point
	gosnmp *g.GoSNMP
	out    chan<- models.Element // TODO add some serious tool for this
}

func (s *SnmpHandler) Handle(pckt []byte) error {
	return s.HandlePacket(pckt)
}

func (snmp *SnmpHandler) HandlePacket(pckt []byte) error {
	trap, err := snmp.gosnmp.UnmarshalTrap(pckt, false) // only supports v2 now
	if err != nil {
		panic(err) // TODO
	}
	p := models.NewPacket(trap)
	b, err := serdes.EncodeJson(&p)
	if err != nil {
		logrus.Info("Encoding failed!!")
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
