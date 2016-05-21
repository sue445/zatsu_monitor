package main

import (
	"bytes"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

type ZatsuMonitor struct {
	databaseFile string
}

const (
	NOT_FOUND_KEY = -1
)

func NewZatsuMonitor(databaseFile string) *ZatsuMonitor {
	z := new(ZatsuMonitor)
	z.databaseFile = databaseFile
	return z
}

func (z ZatsuMonitor) GetDbStatus(key string) (int, error) {
	db, err := leveldb.OpenFile(z.databaseFile, nil)
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

func (z ZatsuMonitor) SaveDbStatus(key string, statusCode int) error {
	db, err := leveldb.OpenFile(z.databaseFile, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutVarint(buf, int64(statusCode))

	db.Put([]byte(key), buf, nil)

	return nil
}
