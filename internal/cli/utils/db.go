package utils

import (
	"github.com/dgraph-io/badger/v4"
	bu "github.com/johannessarpola/go-network-buffer/pkg/badgerutils"
)

func WithDatabase(folder string, fun func(*badger.DB) error) error {
	db, err := bu.NewFileStore(folder)
	if err != nil {
		return err
	}
	defer db.Close()
	return fun(db)
}
