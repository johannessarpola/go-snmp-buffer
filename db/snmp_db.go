package db

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	u "github.com/johannessarpola/go-network-buffer/utils"
)

var _ = u.ConvertToByteArr()
var _ = m.NewIndex("abc", 1)

//var logger = logrus.New() // TODO Fix scope

// TODO Hand multiple offsets?
type SnmpDB struct {
	DB             *badger.DB
	current_idx    m.Index
	current_idx_mu sync.Mutex
	offset_idx     m.Index
	offset_idx_mu  sync.Mutex
	prefix         []byte
}

func (data *SnmpDB) SetCapacity(size uint64) {

}

func (data *SnmpDB) Capacity() uint64 {
	return 0
}

func (data *SnmpDB) Enqueue(v []byte) {

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
