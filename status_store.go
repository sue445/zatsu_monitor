package main

import (
	"bytes"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

type StatusStore struct {
	databaseFile string
}

const (
	NOT_FOUND_KEY = -1
)

func NewStatusStore(databaseFile string) *StatusStore {
	s := new(StatusStore)
	s.databaseFile = databaseFile
	return s
}

func (s *StatusStore) GetDbStatus(key string) (int, error) {
	db, err := leveldb.OpenFile(s.databaseFile, nil)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	if ret, _ := db.Has([]byte(key), nil); !ret {
		return NOT_FOUND_KEY, nil
	}

	data, err := db.Get([]byte(key), nil)

	buf := bytes.NewBuffer(data)
	statusCode, _ := binary.Varint(buf.Bytes())

	return int(statusCode), nil
}

func (s *StatusStore) SaveDbStatus(key string, statusCode int) error {
	db, err := leveldb.OpenFile(s.databaseFile, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutVarint(buf, int64(statusCode))

	db.Put([]byte(key), buf, nil)

	return nil
}
