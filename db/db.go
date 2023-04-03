package db

import (
	"github.com/dgraph-io/badger/v4"
	m "github.com/johannessarpola/go-network-buffer/models"
	u "github.com/johannessarpola/go-network-buffer/utils"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type Data struct {
	DB          *badger.DB
	current_idx m.Index
	offset_idx  m.Index
	prefix      []byte
}

func NewData(path string, prefix string) *Data {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		logger.Fatal("Could not open filestore")
	}

	current_idx := m.ZeroIndex("current_idx")
	offset_idx := m.ZeroIndex("offset_idx")

	d := &Data{
		DB:          db,
		current_idx: current_idx,
		offset_idx:  offset_idx,
		prefix:      []byte(prefix),
	}

	d.init_index(&d.current_idx)
	d.init_index(&d.offset_idx)

	return d

}

func (data *Data) init_index(idx *m.Index) {
	err := data.DB.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(idx.KeyAsBytes())

		if err == nil {
			item.Value(func(val []byte) error {
				n := u.ConvertToUint64(val)
				logger.Infof("Value exists, setting %s from db to %d", idx.Name, n)
				idx.SetValue(n)
				return nil
			})
		} else {
			logger.Infof("Value does not exist, setting %s to 0", idx.Name)
		}

		err = txn.Set(idx.AsBytes())
		return err

	})

	if err != nil {
		logger.Fatalf("Could not initialize index %s", idx.Name)
	}
}

func (data *Data) Append(input []byte) error {
	return data.DB.Update(func(txn *badger.Txn) error {
		_, err := data.IncreaseCurrentIndex()
		if err != nil {
			logger.Info("Could not increase current index")
		}
		txn.Set(data.prefixed_current_idx(data.current_idx.ValueAsBytes()), input)
		return nil
	})
}

func (data *Data) prefixed_current_idx(key []byte) []byte {
	return append(data.prefix, key...)
}

func (data *Data) GetCurrentIndex() uint64 {
	return data.current_idx.Value
}

func (data *Data) GetOffsetIndex() uint64 {
	return data.offset_idx.Value
}

func (data *Data) IncreaseCurrentIndex() (uint64, error) {
	err := data.DB.Update(func(txn *badger.Txn) error {
		data.current_idx.Increment()
		return txn.Set(data.current_idx.AsBytes())
	})
	return data.current_idx.Value, err
}

func (data *Data) Connect(c <-chan []byte) {

	// Debug print
	n := data.GetCurrentIndex()
	logger.Infof("\n%d", n)

	for v := range c {
		data.Append(v)

		// Debug print
		n := data.GetCurrentIndex()
		logger.Info("\n%d", n)
	}

	defer data.DB.Close()
}

// Can be used to move offset forward, save(bool) to control if it is persisted also
func (data *Data) IncrementGetOffset(save bool) (*m.Index, error) {
	data.offset_idx.Increment()
	if save {
		data.UpdateOffset(data.current_idx.Value + 1)
	}
	return &data.offset_idx, nil
}

func (data *Data) UpdateOffset(new_val uint64) error {
	data.offset_idx.SetValue(new_val)
	err := data.SaveIndex(data.offset_idx)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) SaveIndex(idx m.Index) error {
	err := data.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(idx.AsBytes())
	})
	return err

}

func (data *Data) GetIndex(index_name string) (m.Index, error) {
	idx := m.ZeroIndex(index_name) // Have initial struct
	err := data.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(idx.KeyAsBytes())

		if err != nil {
			logger.Info("Index not found")
			return err
		}
		return item.Value(func(val []byte) error {
			idx.SetValue(u.ConvertToUint64(val))
			return nil
		})
	})

	return idx, err
}
