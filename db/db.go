package db

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.ErrorLevel) // TODO Configurable
}

// TODO Clean up the hierarchy after refactoring, this is quite useless class, should access RingDB direct
type Database struct {
	db     *badger.DB
	RingDB *RingDB
	prefix []byte
}

// // TODO Add batch support so it is offset + n
// func (data *Database) GetOffsettedStream(goroutines int, id string) *badger.Stream {
// 	stream := data.db.NewStream()

// 	// -- Optional settings
// 	stream.NumGo = goroutines // Set number of goroutines to use for iteration.
// 	stream.Prefix = data.prefix
// 	stream.LogPrefix = id

// 	// ChooseKey is called concurrently for every key. If left nil, assumes true by default.
// 	stream.ChooseKey = data.after_offset_choosekey

// 	return stream
// }

// func (data *Database) after_offset_choosekey(item *badger.Item) bool {
// 	k := item.Key()[len(data.prefix):]
// 	n := u.ConvertToUint64(k)
// 	o, _ := data.RingDB.IndexDB.oidx_store.GetNbr() // TODO Clean up accessors and handling
// 	return n > o
// }

func NewDatabase(path string, prefix string) *Database {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		logger.Fatal("Could not open filestore")
	}

	rdb := NewRingDB(db, prefix)

	d := &Database{
		db:     db,
		RingDB: rdb,
		prefix: []byte(prefix), // TODO Remove and move to ringdb
	}

	return d

}
