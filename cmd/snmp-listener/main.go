package main

import (
	s "github.com/johannessarpola/go-network-buffer/pkg/snmp"
	sdb "github.com/johannessarpola/go-network-buffer/pkg/snmpdb"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.ErrorLevel)

	port := 9999
	logrus.Infof("Starting snmp listener at port %d", port)
	idx_fn := "../../_idxs"
	snmp_fn := "../../_snmp"
	snmp_pfx := "snmp_"

	snmp_data, err := sdb.OpenDatabases(snmp_fn, snmp_pfx, idx_fn)

	defer snmp_data.Close()
	if err != nil {
		logrus.Fatal("Could not open databases", err)
	}
	s.Start(port, snmp_data)
}
