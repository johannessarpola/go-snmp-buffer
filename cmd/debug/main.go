package main

import (
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	folder := "../../_tmp"
	prefix := "snmp_"
	log.Info("Starting database service on folder %s with prefix %s", folder, prefix)
	data := db.NewDatabase(folder, prefix)

	cidx, _ := data.RingDB.IndexDB.GetCurrentIndex()
	oidx, _ := data.RingDB.IndexDB.GetOffsetIndex()
	log.Infof("Current idx: %d", cidx)
	log.Infof("Offset idx: %d", oidx)
}
