package main

import (
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/listener"
	"github.com/johannessarpola/go-network-buffer/utils"
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

//		// Only log the warning severity or above.
//		log.SetLevel(log.Default())
//	}

func main() {
	var log = logrus.New()

	port := 9999
	log.Infof("Starting snmp listener at port %d", port)
	idx_fs, err := utils.NewFileStore("../../_idxs")
	if err != nil {
		log.Fatalf("could not open index filestore")
	}
	snmp_fs, err := utils.NewFileStore("../../_snmp")
	if err != nil {
		log.Fatalf("could not open snmp filestore")
	}

	defer idx_fs.Close()  // TODO Handle these more cleanly
	defer snmp_fs.Close() // TODO Handle these more cleanly

	idx_db := db.NewIndexDB(idx_fs)                     // TODO prefix?
	snmp_data := db.NewSnmpDB(snmp_fs, idx_db, "snmp_") // TODO Configurable prefix

	listener.Start(port, snmp_data)
}
