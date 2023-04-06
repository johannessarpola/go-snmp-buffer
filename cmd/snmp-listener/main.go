package main

import (
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/listener"
	"github.com/sirupsen/logrus"
)

// TODO Add json logging at some point for cloud native stuffs
// func init() {

// 	if strings.EqualFold(os.Getenv("LOGGING_FORMAT"), "JSON") {
// 		// Log as JSON instead of the default ASCII formatter.
// 		log.SetFormatter(&log.JSONFormatter{})
// 	}

// 	// Output to stdout instead of the default stderr
// 	// Can be any io.Writer, see below for File example
// 	log.SetOutput(os.Stdout)

// 	// Only log the warning severity or above.
// 	log.SetLevel(log.Default())
// }

func main() {
	var log = logrus.New()

	port := 9999
	log.Info("Starting snmp listener at port %s", port)
	folder := "../../_tmp"
	prefix := "snmp_"
	data := db.NewDatabase(folder, prefix)

	log.Info("Starting database service on folder %s with prefix %s", folder, prefix)
	listener.Start(port, data)
}
