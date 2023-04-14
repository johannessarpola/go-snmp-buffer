package db

import (
	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

func IndexFrom(k []byte, v []byte) m.Index {
	return m.NewIndex(string(k), u.ConvertToUint64(v))
}

func ListIndexes(db *badger.DB) []*m.Index {
	iteratorOpts := badger.DefaultIteratorOptions
	iteratorOpts.PrefetchValues = true
	c := make([]*m.Index, 10)
	db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(iteratorOpts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			_ = item.Value(func(val []byte) error {
				idx := IndexFrom(k, val)
				c = append(c, &idx)
				return nil
			})
		}
		return nil // TODO Error handling
	})
	return c
}
