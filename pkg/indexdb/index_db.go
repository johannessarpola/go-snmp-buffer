package indexdb

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/pkg/models"
	"github.com/sirupsen/logrus"
)

// TODO Use array
type IndexDB struct {
	cidx_store *IndexStore // has sync.Mutex
	oidx_store *IndexStore // has sync.Mutex
	common_db  *badger.DB
}

func NewIndexDB(db *badger.DB) (*IndexDB, error) {
	cidx := "current_idx"
	current_idx_store, err := NewIndexStore(cidx, db) // TODO Configurable
	if err != nil {
		logrus.Errorf(fmt.Sprintf("could not initialize %s store", cidx), err)
		return nil, err
	}
	oidx := "offset_idx"
	offset_idx_store, err := NewIndexStore(oidx, db) // TODO Configurable
	if err != nil {
		logrus.Error(fmt.Sprintf("could not initialize %s store", oidx), err)
		return nil, err
	}
	d := &IndexDB{
		cidx_store: current_idx_store,
		oidx_store: offset_idx_store,
		common_db:  db,
	}

	return d, nil
}

func (i *IndexDB) Close() {
	i.common_db.Close()
}

func get_and_increment(idx *IndexStore) (uint64, error) {
	nbr, err := idx.GetNbr()
	_, _ = idx.Increment()
	return nbr, err
}

func (data *IndexDB) GetAndIncrementOffset() (uint64, error) {
	return get_and_increment(data.oidx_store)
}

func (data *IndexDB) GetAndIncrementCurrentIndex() (uint64, error) {
	return get_and_increment(data.cidx_store)
}

func (data *IndexDB) GetCurrentIndex() (uint64, error) {
	return data.cidx_store.GetNbr()
}

func (data *IndexDB) GetOffsetIndex() (uint64, error) {
	return data.oidx_store.GetNbr()
}

func (data *IndexDB) IncrementCurrentIndex() (*m.Index, error) {
	return data.cidx_store.Increment()
}

func (data *IndexDB) IncrementOffsetIndex() (*m.Index, error) {
	return data.oidx_store.Increment()
}

func (data *IndexDB) SetCurrentIndex(newval uint64) error {
	return data.cidx_store.Set(newval)
}
func (data *IndexDB) SetOffsetIndex(newval uint64) error {
	return data.oidx_store.Set(newval)
}
