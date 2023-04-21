package db

import (
	"errors"

	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

func IndexFrom(k []byte, v []byte) m.Index {
	return m.NewIndex(string(k), u.ConvertToUint64(v))
}

func ListIndexes(db *badger.DB) ([]*m.Index, error) {
	iteratorOpts := badger.DefaultIteratorOptions
	iteratorOpts.PrefetchValues = true
	c := make([]*m.Index, 0)
	err := db.View(func(txn *badger.Txn) error {
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
		return nil
	})
	return c, err
}

func GetIndex(db *badger.DB, key []byte) (*m.Index, error) {
	var idx *m.Index
	err := db.Update(func(txn *badger.Txn) error {
		itm, err := txn.Get(key)
		if err != nil {
			return err
		}
		itm.Value(func(val []byte) error {
			idx = &m.Index{
				Name:  string(itm.Key()),
				Value: u.ConvertToUint64(val),
			}
			return nil
		})
		return err
	})
	return idx, err
}

func CreateIndex(db *badger.DB, key []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		i, err := txn.Get(key)
		// it should not override index
		if i == nil && err != nil {
			init := uint64(0)
			return txn.Set(key, u.ConvertToByteArr(init))
		} else if i != nil {
			return errors.New("index exists")
		} else {
			return err
		}
	})
}

func SetIndexUint(db *badger.DB, key []byte, value uint64) error {
	return SetIndex(db, key, u.ConvertToByteArr(value))
}

func SetIndex(db *badger.DB, key []byte, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		i, err := txn.Get(key)
		// Set should not create new index
		if i != nil {
			return txn.Set(key, value)
		}
		return err
	})
}

func DeleteIndex(db *badger.DB, key []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func WithDatabase(folder string, fun func(*badger.DB) error) error {
	db, err := u.NewFileStore(folder)
	if err != nil {
		return err
	}
	defer db.Close()
	return fun(db)
}
