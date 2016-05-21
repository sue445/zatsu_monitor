package main

import (
	"bytes"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

type ZatsuMonitor struct {
	databaseFile string
}

func NewZatsuMonitor(databaseFile string) *ZatsuMonitor {
	z := new(ZatsuMonitor)
	z.databaseFile = databaseFile
	return z
}

func HttpStatusCode(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func (z ZatsuMonitor) GetDbStatus(key string) (int, error) {
	db, err := leveldb.OpenFile(z.databaseFile, nil)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	if ret, _ := db.Has([]byte(key), nil); !ret {
		return 0, nil
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

func IsSuccessfulStatus(statusCode int) bool {
	n := statusCode / 100

	// Successful: 2xx, 3xx
	return n == 2 || n == 3
}
