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
	data := db.NewData(folder, prefix)

	log.Infof("Current idx: %d", data.GetCurrentIndex())
	log.Infof("Offset idx: %d", data.GetOffsetIndex())
}
