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
}

func NewData(path string) *Data {
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
	}

}

func (data *Data) Append(buf []byte) error {
	return data.db.Update(func(txn *badger.Txn) error {
		cur_idx, _ := data.IncreaseCurrentIndex()
		arr := u.ConvertToByteArr(cur_idx)

		// Update the stored idx
		txn.Set(data.current_idx_key, arr)

		// Add the value
		txn.Set(arr, buf)
		return nil
	})
}

func (data *Data) GetCurrentIndex() (uint64, error) {
	var n uint64
	err := data.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(data.current_idx_key)

		if err != nil {
			log.Fatal("No current idx in database ")
		}

		item.Value(func(val []byte) error {
			n = u.ConvertToUint64(val)
			return nil
		})
		return err

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

	err := data.db.Update(func(txn *badger.Txn) error {
		// TODO cleanup
		return txn.Set(data.current_idx_key, u.ConvertToByteArr(uint64(0)))
	})

	if err != nil {
		log.Fatalf("Apuva")
	}

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
