package serdes

import (
	"reflect"
	"testing"

	g "github.com/gosnmp/gosnmp"
	"github.com/johannessarpola/go-network-buffer/pkg/models"
)

// TODO Add benchmark test
func TestJson(t *testing.T) {

	pdu := g.SnmpPDU{
		Name:  ".1.3.6.1.6.3.1.1.4.1.0",
		Type:  g.ObjectIdentifier,
		Value: ".1.3.6.1.6.3.1.1.5.1",
	}

	pb, _ := g.Default.SnmpEncodePacket(
		g.Trap,
		[]g.SnmpPDU{pdu},
		0,
		0,
	)

	p, _ := g.Default.SnmpDecodePacket(pb)

	pckt := models.NewPacket(p)

	js, _ := EncodeJson(&pckt)
	dpckt, _ := DecodeJson(js)

	if !reflect.DeepEqual(pckt, dpckt) {
		t.Error("Packets were not equal")
	}

}

func TestGob(t *testing.T) {

	pdu := g.SnmpPDU{
		Name:  ".1.3.6.1.6.3.1.1.4.1.0",
		Type:  g.ObjectIdentifier,
		Value: ".1.3.6.1.6.3.1.1.5.1",
	}

	pb, _ := g.Default.SnmpEncodePacket(
		g.Trap,
		[]g.SnmpPDU{pdu},
		0,
		0,
	)

	p, _ := g.Default.SnmpDecodePacket(pb)

	pckt := models.NewPacket(p)

	js, _ := EncodeGob(&pckt)
	dpckt, _ := DecodeGob(js)

	if !reflect.DeepEqual(pckt, dpckt) {
		t.Error("Packets were not equal")
	}

}
