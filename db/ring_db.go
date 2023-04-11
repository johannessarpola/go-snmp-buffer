package db

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/utils"
	// u "github.com/johannessarpola/go-network-buffer/utils"
)

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

func (data *RingDB) prefixed_arr(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *RingDB) get_prefixed_element(idx uint64) (*models.Element, error) {
	m := models.EmptyElement()
	err := data.db.View(func(txn *badger.Txn) error {
		barr := utils.ConvertToByteArr(idx)
		k := data.prefixed_arr(barr)
		item, err := txn.Get(k)
		if err != nil {
			logger.Errorf("Could not get element for %d", idx)
		}
		if item != nil {
			item.Value(func(val []byte) error {
				m.Value = val
				return nil
			})
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (data *RingDB) SetCapacity(size uint64) error {
	// TODO Should set the maximunn number of items stored
	data.Lock()
	defer data.Unlock()
	return data.IndexDB.cidx_store.Set(size)
}

func (data *RingDB) ContentSize() (uint64, error) {
	return data.IndexDB.cidx_store.GetNbr()
}

func (data *RingDB) Capacity() (uint64, error) {
	// TODO Should return the number of events stored at max
	return 0, nil
}

func (data *RingDB) Enqueue(v models.Element) error {
	// TODO Add the resetting to 0 when capacity is exceeded
	logger.Info("appending stuff")
	data.Lock()
	idx, err := data.IndexDB.GetAndIncrementCurrentIndex()
	data.Unlock()
	if err != nil {
		logger.Info("Could not increase current index")
	}
	logger.Infof("current idx: %d", idx)
	return data.db.Update(func(txn *badger.Txn) error {

		arr := utils.ConvertToByteArr(idx)
		k := data.prefixed_arr(arr)
		txn.Set(k, v.Value)
		return nil // TODO
	})

}

func (data *RingDB) Dequeue() (*models.Element, error) {
	// TODO Fix issue where this allows offset > current idx
	// you could also delete it here as according to spec but maybe later, or add a cleanup job on separate cmd
	logger.Info("dequeue stuff")
	idx, err := data.IndexDB.GetAndIncrementOffset()
	logger.Infof("current offset: %d", idx)
	if err != nil {
		logger.Info("Could not increase offset index")
	}
	cs, _ := data.ContentSize()
	if idx == cs {
		return nil, nil
	} else {
		return data.get_prefixed_element(idx)
	}

}

func (data *RingDB) Peek() (*models.Element, error) {
	logger.Info("peeking stuff")
	data.Lock()
	oidx, err := data.IndexDB.oidx_store.GetNbr()
	data.Unlock()
	if err != nil {
		logger.Error("Could not get offset index")
	}

	return data.get_prefixed_element(oidx)
}

func (data *RingDB) Values() []*models.Element {
	// TODO Return stream
	return nil
}
