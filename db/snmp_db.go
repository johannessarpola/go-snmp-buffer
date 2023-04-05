package db

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	// u "github.com/johannessarpola/go-network-buffer/utils"
)

var _ = m.NewIndex("abc", 1)

//var logger = logrus.New() // TODO Fix scope

// TODO Hand multiple offsets?
type SnmpDB struct {
	sync.Mutex
	DB      *badger.DB
	IndexDB *IndexDB
	prefix  []byte
}

func (data *SnmpDB) prefixed_current_idx(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *SnmpDB) SetCapacity(size uint64) {
}

func (data *SnmpDB) Capacity() uint64 {
	return 0
}

func (data *SnmpDB) Enqueue(v []byte) error {
	data.Lock()
	defer data.Unlock()

	logger.Info("appending event")
	return data.DB.Update(func(txn *badger.Txn) error {
		idx, err := data.IndexDB.IncrementCurrentIndex()
		logger.Infof("current idx: %d", idx)
		if err != nil {
			logger.Info("Could not increase current index")
		}
		k := data.prefixed_current_idx(idx.ValueAsBytes())
		txn.Set(k, v)
		return nil
	})

}

func (data *SnmpDB) Dequeue() []byte {
	return nil
}

func (data *SnmpDB) Peek() []byte {
	return nil
}

func (data *SnmpDB) Values() [][]byte {
	return nil
}
