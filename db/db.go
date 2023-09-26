package db

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LevelDB struct {
	db *leveldb.DB
}

func NewLevelDB(filename string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(filename, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDB{db}, nil
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	key = checkNilBytes(key)
	res, err := db.db.Get(key, nil)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (db *LevelDB) Set(key []byte, value []byte) error {
	key = checkNilBytes(key)
	value = checkNilBytes(value)
	err := db.db.Put(key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db *LevelDB) Delete(key []byte) error {
	key = checkNilBytes(key)
	err := db.db.Delete(key, nil)
	if err != nil {
		log.Println("ERROR DELETING DB KEY:", err)
		return err
	}
	return nil
}

func (db *LevelDB) GetAll(prefix string) [][]byte {
	iter := db.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	values := make([][]byte, 0)
	for iter.Next() {
		values = append(values, iter.Value())
	}
	iter.Release()
	return values
}
