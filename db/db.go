package db

import (
	"log"
	"strconv"

	"github.com/dgraph-io/badger/v3"
)

type Data struct{}

func (data *Data) Start(path string, c <-chan []byte) {
	opts := badger.DefaultOptions(path)

	db, err := badger.Open(opts) // TODO Define path better sometime
	if err != nil {
		log.Fatal(err)
	}

	last_key := []byte("last")
	last_val := []byte(strconv.Itoa(0))
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set(last_key, last_val)
	})

	if err != nil {
		log.Fatalf("Apuva")
	}

	// Debug print
	db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get(last_key)
		value := make([]byte, item.ValueSize())
		item.ValueCopy(value)
		last_val, _ := strconv.Atoi(string(value))
		println(last_val)
		return nil
	})

	for v := range c {
		// TODO Needs lots of cleaning up
		_ = db.Update(func(txn *badger.Txn) error {
			item, _ := txn.Get(last_key)
			value := make([]byte, item.ValueSize())
			item.ValueCopy(value)

			lastInt, _ := strconv.Atoi(string(value))
			lastInt++
			println(lastInt)
			txn.Set(last_key, []byte(strconv.Itoa(lastInt)))
			txn.Set(last_val, v)
			return nil
		})

		// Debug print
		db.View(func(txn *badger.Txn) error {
			item, _ := txn.Get(last_key)
			value := make([]byte, item.ValueSize())
			item.ValueCopy(value)
			last_val, _ := strconv.Atoi(string(value))
			println(last_val)
			return nil
		})

		println("Recv")
	}

	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("Key"), []byte("Value"))
	})

}
