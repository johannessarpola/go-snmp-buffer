package snmpdb

import (
	"github.com/dgraph-io/badger/v4"
	_ "github.com/johannessarpola/go-network-buffer/pkg/logging"
	m "github.com/johannessarpola/go-network-buffer/pkg/models"
	s "github.com/johannessarpola/go-network-buffer/pkg/serdes"
	"github.com/sirupsen/logrus"
)

func LastN(db *badger.DB, dst []m.StoredPacket) error {

	n := len(dst)
	i := 0

	iteratorOpts := badger.DefaultIteratorOptions
	iteratorOpts.Reverse = true
	iteratorOpts.PrefetchValues = true

	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(iteratorOpts)
		defer it.Close()
		for it.Rewind(); it.Valid() && i < n; it.Next() {
			item := it.Item()
			k := item.Key()
			logrus.Debugf("Listing trap with id %s", string(k))
			_ = item.Value(func(val []byte) error {
				decoded, err := s.DecodeJson(val)

				if err == nil {
					dst[i] = m.StoredPacket{
						Key:    string(k),
						Packet: decoded,
					}
				}
				i++

				return err
			})
		}
		return nil
	})
	return err
}
