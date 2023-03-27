package db

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

type Data struct {
	db              *badger.DB
	current_idx_key []byte // BigEndian, uint64 // TODO Should use sequences from badgerdb?
	offset_idx_key  []byte // BigEndian, uint64 // TODO Should use sequences from badgerdb?
	prefix          []byte
}

func NewData(path string, prefix string) *Data {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal("Could not open filestore")
	}

	current_idx_key := []byte("current_idx")
	offset_idx_key := []byte("offset_idx")

	return &Data{
		db:              db,
		current_idx_key: current_idx_key, //
		offset_idx_key:  offset_idx_key,
		prefix:          []byte(prefix),
	}

}

func (data *Data) Append(input []byte) error {
	return data.db.Update(func(txn *badger.Txn) error {
		cur_idx, _ := data.IncreaseCurrentIndex()
		key := u.ConvertToByteArr(cur_idx)
		// Update the stored idx
		txn.Set(data.current_idx_key, key)

		// Add the value
		txn.Set(data.prefixed_current_idx(key), input)
		return nil
	})
}

func (data *Data) prefixed_current_idx(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *Data) GetCurrentIndex() (uint64, error) {
	var n uint64 = 0 // TODO Add a flag to notify if it is empty
	err := data.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(data.current_idx_key)

		if err == nil {
			item.Value(func(val []byte) error {
				println("Value exists, setting cur_idx from db")
				n = u.ConvertToUint64(val)
				return nil
			})
		} else {
			println("Value does not exist, setting cur_idx to 0")
			txn.Set(data.current_idx_key, u.ConvertToByteArr(n))
		}

		return nil // TODO Fix

	})
	return n, err
}

func (data *Data) IncreaseCurrentIndex() (uint64, error) {
	var r uint64
	err := data.db.Update(func(txn *badger.Txn) error {
		n, err := data.GetCurrentIndex()
		r = n + 1
		if err != nil {
			log.Fatal("No current idx in database ")
		}

		return txn.Set(data.current_idx_key, u.ConvertToByteArr(r))
	})
	return r, err
}

func (data *Data) Connect(c <-chan []byte) {

	// Debug print
	n, _ := data.GetCurrentIndex()
	fmt.Printf("\n%d", n)

	for v := range c {
		// TODO Needs lots of cleaning up
		data.Append(v)

		// Debug print
		n, _ := data.GetCurrentIndex()
		fmt.Printf("\n%d", n)
	}

	defer data.db.Close()
}
