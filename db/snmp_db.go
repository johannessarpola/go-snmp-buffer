package db

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.ErrorLevel) // TODO Configurable
}

type SnmpDB struct {
	Buffer *RingDB
}

func NewSnmpDB(fs *badger.DB, idx_db *IndexDB, prefix string) *SnmpDB {
	rdb := NewRingDB(fs, idx_db, prefix)

	d := &SnmpDB{
		Buffer: rdb,
	}

	return d

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
