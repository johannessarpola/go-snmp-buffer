package main

import (
	bu "github.com/johannessarpola/go-network-buffer/pkg/badgerutils"
	idb "github.com/johannessarpola/go-network-buffer/pkg/indexdb"
	s "github.com/johannessarpola/go-network-buffer/pkg/snmp"
	sdb "github.com/johannessarpola/go-network-buffer/pkg/snmpdb"
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
	logrus.SetLevel(logrus.ErrorLevel)

	port := 9999
	logrus.Infof("Starting snmp listener at port %d", port)
	idx_fs, err := bu.NewFileStore("../../_idxs")
	if err != nil {
		logrus.Fatalf("could not open index filestore")
	}
	snmp_fs, err := bu.NewFileStore("../../_snmp")
	if err != nil {
		logrus.Fatalf("could not open snmp filestore")
	}

	defer idx_fs.Close()  // TODO Handle these more cleanly
	defer snmp_fs.Close() // TODO Handle these more cleanly

	idx_db := idb.NewIndexDB(idx_fs)                     // TODO prefix?
	snmp_data := sdb.NewSnmpDB(snmp_fs, idx_db, "snmp_") // TODO Configurable prefix

	s.Start(port, snmp_data)
}
