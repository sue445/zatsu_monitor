package main

import (
	"bytes"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

var OK_STATUS_CODES = []int{200, 301, 302, 303, 304, 307, 308}

const (
	STATUS_TRUE  = "1"
	STATUS_FALSE = "0"
)

type ZatsuMonitor struct {
	databaseFile string
}

func NewZatsuMonitor(databaseFile string) *ZatsuMonitor {
	z := new(ZatsuMonitor)
	z.databaseFile = databaseFile
	return z
}

func (z ZatsuMonitor) CheckUrl(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		return false
	}

	for _, v := range OK_STATUS_CODES {
		if v == resp.StatusCode {
			return true
		}
	}

	return false
}

func (z ZatsuMonitor) GetDbStatus(key string) bool {
	db, err := leveldb.OpenFile(z.databaseFile, nil)
	if err != nil {
		panic("Failed: OpenFile " + z.databaseFile)
	}
	defer db.Close()

	if ret, _ := db.Has([]byte(key), nil); !ret {
		// If not exists key, same as success (1st operation)
		return true
	}

	data, err := db.Get([]byte(key), nil)

	buf := bytes.NewBuffer(data)
	value := buf.String()

	switch value {
	case STATUS_TRUE:
		return true
	case STATUS_FALSE:
		return false
	default:
		panic("Unknown value: " + value)
	}
}

func (z ZatsuMonitor) SaveDbStatus(key string, value bool) {
	db, err := leveldb.OpenFile(z.databaseFile, nil)
	if err != nil {
		panic("Failed: OpenFile " + z.databaseFile)
	}
	defer db.Close()

	var data string
	if value {
		data = STATUS_TRUE
	} else {
		data = STATUS_FALSE
	}
	db.Put([]byte(key), []byte(data), nil)
}
