package db

import (
	"log"
	"strconv"

	"github.com/dgraph-io/badger/v3"
)

type Data struct{}

func (data *Data) Start(path string, c chan []byte) {
	db, err := badger.Open(badger.DefaultOptions(path)) // TODO Define path better sometime
	if err != nil {
		log.Fatal(err)
	}

	//first := badger.NewEntry([]byte("first"), []byte(strconv.Itoa(0)))
	last := badger.NewEntry([]byte("last"), []byte(strconv.Itoa(0)))

	for v := range c {
		err = db.Update(func(txn *badger.Txn) error {
			item, _ := txn.Get(last.Key)
			value := make([]byte, item.ValueSize())
			item.ValueCopy(value)

			lastInt, _ := strconv.Atoi(string(value))
			lastInt++
			println(lastInt)
			txn.Set(last.Key, []byte(strconv.Itoa(lastInt)))
			txn.Set(last.Value, v)
			return nil
		})

		// Debug print
		db.View(func(txn *badger.Txn) error {
			item, _ := txn.Get(last.Key)
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
