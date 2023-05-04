package db

import (
	"github.com/dgraph-io/badger/v4"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

func WithDatabase(folder string, fun func(*badger.DB) error) error {
	db, err := u.NewFileStore(folder)
	if err != nil {
		return err
	}
	defer db.Close()
	return fun(db)
}
