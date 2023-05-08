package badgerutils

import (
	"github.com/dgraph-io/badger/v4"
)

func NewFileStore(folder string) (*badger.DB, error) {
	// TODO Handle opts at some point
	opts := badger.DefaultOptions(folder)
	opts = opts.WithLoggingLevel(badger.ERROR) // TODO Handle logging level somehow
	return badger.Open(opts)
}
