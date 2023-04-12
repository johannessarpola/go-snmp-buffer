package main

import (
	"fmt"
	"log"
	"os"
	"time"

	g "github.com/gosnmp/gosnmp"
)

// TODO use args to count and some variation
func main() {

	params := &g.GoSNMP{
		Target:        "127.0.0.1",
		Port:          9999,
		Version:       g.Version3,
		Timeout:       time.Duration(30) * time.Second,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Logger:        g.NewLogger(log.New(os.Stdout, "", 0)),
		SecurityParameters: &g.UsmSecurityParameters{UserName: "user",
			AuthoritativeEngineID:    "1234",
			AuthenticationProtocol:   g.SHA,
			AuthenticationPassphrase: "password",
			PrivacyProtocol:          g.DES,
			PrivacyPassphrase:        "password",
		},
	}
	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	// TODO Mutation
	pdu := g.SnmpPDU{
		Name:  ".1.3.6.1.6.3.1.1.4.1.0",
		Type:  g.ObjectIdentifier,
		Value: ".1.3.6.1.6.3.1.1.5.1",
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{pdu},
	}

	for i := 0; i <= 1000000; i++ { // TODO from arg
		_, err = params.SendTrap(trap)
		fmt.Printf("Send trap nbr %d\n", i)
		if err != nil {
			panic(err)
		}
	}

}
