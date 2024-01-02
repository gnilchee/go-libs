package tinyds

import (
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

var (
	DefaultOptions = badger.DefaultOptions
)

type TinyDS struct {
	db *badger.DB
}

func Open(opts badger.Options) *TinyDS {
	tds := TinyDS{}

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("DB create or open failed: %s\n", err)
	}
	tds.db = db
	return &tds
}

func (tds *TinyDS) Set(key string, val string) error {
	var err error
	err = tds.db.Update(func(txn *badger.Txn) error {
		err = txn.Set([]byte(key), []byte(val))
		return err
	})
	return err
}

func (tds *TinyDS) SetwithTTL(key string, val string, ttl int) error {
	var err error
	err = tds.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(val)).WithTTL(time.Duration(ttl) * time.Second)
		err = txn.SetEntry(e)
		return err
	})
	return err
}

func (tds *TinyDS) Get(key string) (string, error) {
	var val []byte
	err := tds.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val_err := item.Value(func(kval []byte) error {
			val = kval
			return nil
		})
		if val_err != nil {
			return val_err
		}
		return nil
	})
	return string(val), err
}

func (tds *TinyDS) Delete(key string) error {
	var err error
	err = tds.db.Update(func(txn *badger.Txn) error {
		err = txn.Delete([]byte(key))
		return err
	})
	return err
}

func (tds *TinyDS) Close() error {
	return tds.db.Close()
}
