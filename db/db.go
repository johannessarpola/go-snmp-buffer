package db

import (
	"log"
	"strconv"

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
		item, _ := txn.Get(data.current_idx_key)
		value := make([]byte, item.ValueSize())
		item.ValueCopy(value)

		lastInt, _ := strconv.Atoi(string(value))
		lastInt++
		txn.Set(data.current_idx_key, []byte(strconv.Itoa(lastInt)))
		txn.Set(last_val, v)
		return nil
	})
}

func (data *Data) IncreaseCurrentIndex() {
	data.db.Update(func(txn *badger.Txn) error {
		v, err := txn.Get(db.current_idx_key)

		if err != nil {
			log.Fatal("No current idx in database ")
		}

		err = v.Value(func(val []byte) error {
			n := u.ConvertToUint64(val)
			n++

			return txn.Set(data.current_idx_key, u.ConvertToByteArr(n))
		})

		if err != nil {
			log.Fatal("Could not update current idx value in database")
		}

	})
}

func (data *Data) Start(c <-chan []byte) {

	last_val := []byte(strconv.Itoa(0))
	err := data.db.Update(func(txn *badger.Txn) error {
		return txn.Set(data.current_idx_key, last_val)
	})

	if err != nil {
		log.Fatalf("Apuva")
	}

	// Debug print
	data.db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get(data.current_idx_key)
		value := make([]byte, item.ValueSize())
		item.ValueCopy(value)
		last_val, _ := strconv.Atoi(string(value))
		println(last_val)
		return nil
	})

	for v := range c {
		// TODO Needs lots of cleaning up
		_ = data.db.Update(func(txn *badger.Txn) error {
			item, _ := txn.Get(data.current_idx_key)
			value := make([]byte, item.ValueSize())
			item.ValueCopy(value)

			lastInt, _ := strconv.Atoi(string(value))
			lastInt++
			println(lastInt)
			txn.Set(data.current_idx_key, []byte(strconv.Itoa(lastInt)))
			txn.Set(last_val, v)
			return nil
		})

		// Debug print
		data.db.View(func(txn *badger.Txn) error {
			item, _ := txn.Get(data.current_idx_key)
			value := make([]byte, item.ValueSize())
			item.ValueCopy(value)
			last_val, _ := strconv.Atoi(string(value))
			println(last_val)
			return nil
		})

		println("Recvd %s", string(v))
	}

	defer data.db.Close()
}
