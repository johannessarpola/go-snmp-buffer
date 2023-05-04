package db

import (
	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	s "github.com/johannessarpola/go-network-buffer/serdes"
)

func LastN(db *badger.DB, dst []m.Packet) error {

	n := len(dst)
	i := 0

	iteratorOpts := badger.DefaultIteratorOptions
	iteratorOpts.Reverse = true
	iteratorOpts.PrefetchValues = true

	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(iteratorOpts)
		defer it.Close()
		for it.Rewind(); it.Valid() && i < n; it.Next() {
			item := it.Item()
			k := item.Key()
			logger.Debugf("Listing trap with id %s", string(k))
			_ = item.Value(func(val []byte) error {
				decoded, err := s.Decode(val)

				if err == nil {
					dst[i] = decoded
				}
				i++

				return err
			})
		}
		return nil
	})
	return err
}

func Last10(db *badger.DB) ([]m.Packet, error) {
	arr := make([]m.Packet, 10)
	return arr, LastN(db, arr)
}
