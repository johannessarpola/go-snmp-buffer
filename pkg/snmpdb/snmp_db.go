package snmpdb

import (
	"sync"

	"github.com/dgraph-io/badger/v4"
	bu "github.com/johannessarpola/go-network-buffer/pkg/badgerutils"
	c "github.com/johannessarpola/go-network-buffer/pkg/conversions"
	i "github.com/johannessarpola/go-network-buffer/pkg/indexdb"
	"github.com/johannessarpola/go-network-buffer/pkg/models"
	"github.com/sirupsen/logrus"
)

// TODO Hand multiple offsets?
type ringDB struct {
	sync.Mutex
	db      *badger.DB
	IndexDB *i.IndexDB // TODO Probably does not need to be accesible from outside
	prefix  []byte
}

func NewRingDB(fs *badger.DB, idx_db *i.IndexDB, prefix string) *ringDB {
	return &ringDB{
		db:      fs,
		IndexDB: idx_db,
		prefix:  []byte(prefix),
	}
}

func (data *ringDB) prefixed_arr(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *ringDB) get_prefixed_element(idx uint64) (*models.Element, error) {
	m := models.EmptyElement()
	err := data.db.View(func(txn *badger.Txn) error {
		barr := c.ConvertToByteArr(idx)
		k := data.prefixed_arr(barr)
		item, err := txn.Get(k)
		if err != nil {
			logrus.Errorf("Could not get element for %d", idx)
		}
		if item != nil {
			err = item.Value(func(val []byte) error {
				m.Value = val
				return nil
			})
			if err != nil {
				logrus.Error("Could not set value")
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (data *ringDB) SetCapacity(size uint64) error {
	// TODO Should set the maximunn number of items stored
	data.Lock()
	defer data.Unlock()
	return data.IndexDB.SetCurrentIndex(size)
}

func (data *ringDB) ContentSize() (uint64, error) {
	return data.IndexDB.GetCurrentIndex()
}

func (data *ringDB) Capacity() (uint64, error) {
	// TODO Should return the number of events stored at max
	return 0, nil
}

func (data *ringDB) Enqueue(v models.Element) error {
	// TODO Add the resetting to 0 when capacity is exceeded
	logrus.Info("appending stuff")
	data.Lock()
	idx, err := data.IndexDB.GetAndIncrementCurrentIndex()
	data.Unlock()
	if err != nil {
		logrus.Info("Could not increase current index")
	}
	logrus.Infof("current idx: %d", idx)
	return data.db.Update(func(txn *badger.Txn) error {

		arr := c.ConvertToByteArr(idx)
		k := data.prefixed_arr(arr)
		err := txn.Set(k, v.Value)
		if err != nil {
			logrus.Error("Could not set value", err)
		}
		return err
	})

}

func (data *ringDB) Dequeue() (*models.Element, error) {
	// TODO Fix issue where this allows offset > current idx
	// you could also delete it here as according to spec but maybe later, or add a cleanup job on separate cmd
	logrus.Info("dequeue stuff")
	idx, err := data.IndexDB.GetAndIncrementOffset()
	logrus.Infof("current offset: %d", idx)
	if err != nil {
		logrus.Info("Could not increase offset index")
	}
	cs, _ := data.ContentSize()
	if idx == cs {
		return nil, nil
	} else {
		return data.get_prefixed_element(idx)
	}

}

func (data *ringDB) Peek() (*models.Element, error) {
	logrus.Info("peeking stuff")
	data.Lock()
	oidx, err := data.IndexDB.GetOffsetIndex()
	data.Unlock()
	if err != nil {
		logrus.Error("Could not get offset index")
	}

	return data.get_prefixed_element(oidx)
}

func (data *ringDB) Values() []*models.Element {
	// TODO Return stream
	return nil
}

type SnmpDB = ringDB

func (data *ringDB) Close() {
	data.db.Close()
	data.IndexDB.Close()
}

func OpenDatabases(snmp_folder string, snmp_prefix string, idx_folder string) (*SnmpDB, error) {
	idx_fs, err := bu.NewFileStore(idx_folder)
	if err != nil {
		logrus.Error("could not open index filestore", err)
		return nil, err
	}
	snmp_fs, err := bu.NewFileStore(snmp_folder)
	if err != nil {
		logrus.Error("could not open snmp filestore", err)
		return nil, err
	}

	idx_db, err := i.NewIndexDB(idx_fs)
	if err != nil {
		return nil, err
	}

	snmp_data := NewRingDB(snmp_fs, idx_db, snmp_prefix)
	return snmp_data, nil
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
