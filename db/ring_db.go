package db

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/johannessarpola/go-network-buffer/models"
	m "github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/utils"
	// u "github.com/johannessarpola/go-network-buffer/utils"
)

var _ = m.NewIndex("abc", 1)

//var logger = logrus.New() // TODO Fix scope

// TODO Hand multiple offsets?
type RingDB struct {
	sync.Mutex
	db      *badger.DB
	IndexDB *IndexDB // TODO Probably does not need to be accesible from outside
	prefix  []byte
}

func NewRingDB(db *badger.DB, prefix string) *RingDB {
	store := NewIndexDB(db)

	return &RingDB{
		db:      db,
		IndexDB: store,
		prefix:  []byte(prefix),
	}
}

func (data *RingDB) prefixed_current_idx(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *RingDB) get_prefixed_element(idx uint64) (*models.Element, error) {
	m := models.EmptyElement()
	err := data.db.View(func(txn *badger.Txn) error {
		barr := utils.ConvertToByteArr(idx)
		item, err := txn.Get(data.prefixed_current_idx(barr))
		if err != nil {
			logger.Errorf("Could not get element for %d", idx)
		}
		m.Value, err = item.ValueCopy(nil) // According to docs should return new array with nil passed
		return err
	})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (data *RingDB) SetCapacity(size uint64) {
}

func (data *RingDB) Capacity() uint64 {
	return 0
}

func (data *RingDB) Enqueue(v models.Element) error {
	data.Lock()
	defer data.Unlock()

	logger.Info("appending stuff")
	idx, err := data.IndexDB.IncrementCurrentIndex()
	logger.Infof("current idx: %d", idx)
	return data.db.Update(func(txn *badger.Txn) error {
		if err != nil {
			logger.Info("Could not increase current index")
		}
		k := data.prefixed_current_idx(idx.ValueAsBytes())
		txn.Set(k, v.Value)
		return nil
	})

}

func (data *RingDB) Dequeue() (*models.Element, error) {
	data.Lock()
	defer data.Unlock()

	logger.Info("dequeue stuff")
	el, err := data.Peek()
	logger.Info("incrementing oidx")
	data.IndexDB.oidx_store.Increment() // As this moves forward its fine to just increase offset
	// you could also delete it here as according to spec but maybe later, or add a cleanup job on separate cmd
	return el, err
}

func (data *RingDB) Peek() (*models.Element, error) {
	logger.Info("peeking stuff")
	oidx, err := data.IndexDB.oidx_store.GetNbr()
	if err != nil {
		logger.Error("Could not get offset index")
	}

	return data.get_prefixed_element(oidx)
}

func (data *RingDB) Values() []*models.Element {
	return nil
}
