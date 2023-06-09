package indexdb

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	c "github.com/johannessarpola/go-network-buffer/pkg/conversions"
	_ "github.com/johannessarpola/go-network-buffer/pkg/logging"
	m "github.com/johannessarpola/go-network-buffer/pkg/models"
	"github.com/sirupsen/logrus"
)

// TODO Could be just IndexStore -> Index with ID
type IndexStore struct {
	sync.Mutex
	db  *badger.DB
	idx *m.Index
}

func NewIndexStore(key string, db *badger.DB) (*IndexStore, error) {
	idx := m.ZeroIndex(key) // TODO More configurable

	d := &IndexStore{
		db:  db,
		idx: &idx,
	}
	err := d.init_index(&idx)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (store *IndexStore) init_index(idx *m.Index) error {
	err := store.db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(idx.KeyAsBytes())

		if err == nil {
			err = item.Value(func(val []byte) error {
				n := c.ConvertToUint64(val)
				logrus.Infof("Value exists, setting %s from db to %d", idx.Name, n)
				idx.SetValue(n)
				return nil
			})
			if err != nil {
				logrus.Error("Could not set value", err)
				return err
			}
		} else {
			logrus.Infof("Value does not exist, setting %s to 0", idx.Name)
		}

		err = txn.Set(idx.AsBytes())
		return err

	})
	return err
}

func (data *IndexStore) GetNbr() (uint64, error) {
	if idx, err := data.Get(); err != nil {
		return 0, err // Zero is the correct vaue when index undefined in this case
	} else {
		return idx.Value, nil
	}
}

func (data *IndexStore) Set(newval uint64) error {
	data.Lock()
	defer data.Unlock()

	return data.db.Update(func(txn *badger.Txn) error {
		return txn.Set(data.idx.KeyAsBytes(), c.ConvertToByteArr(newval))
	})
}

func (data *IndexStore) Increment() (*m.Index, error) {
	data.Lock()
	defer data.Unlock()

	err := data.db.Update(func(txn *badger.Txn) error {
		data.idx.Value = data.idx.Value + 1
		return data.save()
	})
	if err != nil {
		logrus.Error("Error error")
		panic(err) // TODO
	}
	return data.idx, nil
}

func (data *IndexStore) Get() (*m.Index, error) {
	// TODO Could be optimized to be something like withIndex(func(idx )-> T) [single query]
	err := data.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(data.idx.KeyAsBytes())

		if err != nil {
			logrus.Info("Index not found", err)
		}

		return item.Value(func(val []byte) error {
			// Set the structs value from the one in db
			data.idx.SetValue(c.ConvertToUint64(val))
			return nil
		})
	})

	return data.idx, err
}

func (data *IndexStore) save() error {
	err := data.db.Update(func(txn *badger.Txn) error {
		return txn.Set(data.idx.AsBytes())
	})
	return err
}
