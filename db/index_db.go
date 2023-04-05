package db

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

type IndexDB struct {
	DB             *badger.DB
	current_idx    m.Index
	current_idx_mu sync.Mutex
	offset_idx     m.Index
	offset_idx_mu  sync.Mutex
}

func NewIndexDB(path string, db *badger.DB) *IndexDB {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		logger.Fatal("Could not open filestore")
	}

	current_idx := m.ZeroIndex("current_idx") // TODO More configurable
	offset_idx := m.ZeroIndex("offset_idx")   // TODO More configurable

	d := &IndexDB{
		DB:          db,
		current_idx: current_idx,
		offset_idx:  offset_idx,
	}

	d.init_index(&d.current_idx)
	d.init_index(&d.offset_idx)

	return d
}

func (data *IndexDB) init_index(idx *m.Index) {
	err := data.DB.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(idx.KeyAsBytes())

		if err == nil {
			item.Value(func(val []byte) error {
				n := u.ConvertToUint64(val)
				logger.Infof("Value exists, setting %s from db to %d", idx.Name, n)
				idx.SetValue(n)
				return nil
			})
		} else {
			logger.Infof("Value does not exist, setting %s to 0", idx.Name)
		}

		err = txn.Set(idx.AsBytes())
		return err

	})

	if err != nil {
		logger.Fatalf("Could not initialize index %s", idx.Name)
	}
}

func (data *IndexDB) GetCurrentIndex() (uint64, error) {
	return data.GetIndexNbr(data.current_idx.Name)
}

func (data *IndexDB) GetOffsetIndex() (uint64, error) {
	return data.GetIndexNbr(data.offset_idx.Name)
}

func (data *IndexDB) GetIndexNbr(index_name string) (uint64, error) {
	if idx, err := data.GetIndex(index_name); err != nil {
		return 0, err // Zero is the correct vaue when index undefined in this case
	} else {
		return idx.Value, nil
	}
}

func (data *IndexDB) IncrementIndex(index_name string) (*m.Index, error) {
	idx, err := data.GetIndex(index_name)
	if err != nil {
		panic(err) // TODO
	}

	idx.Lock()
	defer idx.Unlock()

	err = data.DB.Update(func(txn *badger.Txn) error {
		idx.Increment()
		return data.SaveIndex(&idx)
	})
	if err != nil {
		panic(err) // TODO
	}
	return &data.current_idx, err
}

func (data *IndexDB) GetIndex(index_name string) (m.Index, error) {
	idx := m.ZeroIndex(index_name) // Have initial struct, 0 is correct for uninitialized
	err := data.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(idx.KeyAsBytes())

		if err != nil {
			logger.Info("Index not found", err)
		}

		return item.Value(func(val []byte) error {
			// Set the structs value from the one in db
			idx.SetValue(u.ConvertToUint64(val))
			return nil
		})
	})

	return idx, err
}

func (data *IndexDB) SaveIndex(idx *m.Index) error {
	idx.Lock()
	defer idx.Unlock()

	err := data.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(idx.AsBytes())
	})
	return err
}
