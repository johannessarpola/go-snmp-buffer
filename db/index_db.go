package db

import (
	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
)

// TODO Could be just IndexDB -> Index with ID
type IndexDB struct {
	db         *badger.DB
	cidx_store *IndexStore
	oidx_store *IndexStore
}

func NewIndexDB(db *badger.DB) *IndexDB {
	current_idx_store := NewIndexStore("current_idx", db) // TODO Configurable
	offset_idx_store := NewIndexStore("offset_idx", db)   // TODO Configurable

	d := &IndexDB{
		db:         db,
		cidx_store: current_idx_store,
		oidx_store: offset_idx_store,
	}

	return d
}

func get_and_increment(idx *IndexStore) (uint64, error) {
	nbr, err := idx.GetNbr()
	idx.Increment()
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
