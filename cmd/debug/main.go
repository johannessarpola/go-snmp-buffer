package main

import (
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()

	idx_fs, err := utils.NewFileStore("../../_idxs")
	if err != nil {
		log.Fatal("could not open index filestore")
	}

	idx_db := db.NewIndexDB(idx_fs)

	cidx, _ := idx_db.GetCurrentIndex()
	oidx, _ := idx_db.GetOffsetIndex()
	log.Infof("Current idx: %d", cidx)
	log.Infof("Offset idx: %d", oidx)
}
