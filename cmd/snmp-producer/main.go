package main

import (
	"fmt"
	g "github.com/gosnmp/gosnmp"
)

// TODO use args to count and some variation
func main() {

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "127.0.0.1"
	g.Default.Port = 9999
	g.Default.Version = g.Version2c
	g.Default.Community = "public"
	//g.Default.Logger = g.NewLogger(log.New(os.Stdout, "", 0))
	err := g.Default.Connect()
	if err != nil {
		panic(err)
		//log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	pdu := g.SnmpPDU{
		Name:  ".1.3.6.1.6.3.1.1.4.1.0",
		Type:  g.ObjectIdentifier,
		Value: ".1.3.6.1.6.3.1.1.5.1",
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{pdu},
	}

	for i := 0; i <= 10000000; i++ { // TODO from arg
		_, err = g.Default.SendTrap(trap)
		fmt.Printf("Send trap nbr %d\n", i)
		if err != nil {
			panic(err)
		}
	}

}
