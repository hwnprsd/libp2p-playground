package db

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type DB struct {
	db *leveldb.DB
}

func NewDB(filename string) (*DB, error) {
	db, err := leveldb.OpenFile(filename, nil)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Get(key []byte) []byte {
	key = checkNilBytes(key)
	res, err := db.db.Get(key, nil)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return res
}

func (db *DB) Set(key []byte, value []byte) {
	key = checkNilBytes(key)
	value = checkNilBytes(value)
	err := db.db.Put(key, value, nil)
	if err != nil {
		log.Println("ERROR SETTING DB KEY:", err)
	}
}

func (db *DB) Delete(key []byte) {
	key = checkNilBytes(key)
	err := db.db.Delete(key, nil)
	if err != nil {
		log.Println("ERROR DELETING DB KEY:", err)
	}
}
