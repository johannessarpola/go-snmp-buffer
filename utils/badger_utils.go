package utils

import (
	"github.com/dgraph-io/badger/v4"
)

func NewFileStore(folder string) (*badger.DB, error) {
	// TODO Handle opts at some point
	opts := badger.DefaultOptions(folder)
	return badger.Open(opts)
}
