package main

import (
	"bytes"
	"encoding/binary"
	"github.com/cockroachdb/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

// StatusStore represents current status cache store
type StatusStore struct {
	databaseFile string
}

const (
	// NotFoundKey represents value if key is not found
	NotFoundKey = -1
)

// NewStatusStore create new StatusStore instance
func NewStatusStore(databaseFile string) *StatusStore {
	s := new(StatusStore)
	s.databaseFile = databaseFile
	return s
}

// GetDbStatus returns status code for specified key
func (s *StatusStore) GetDbStatus(key string) (int, error) {
	db, err := leveldb.OpenFile(s.databaseFile, nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer db.Close()

	if ret, _ := db.Has([]byte(key), nil); !ret {
		return NotFoundKey, nil
	}

	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	buf := bytes.NewBuffer(data)
	statusCode, _ := binary.Varint(buf.Bytes())

	return int(statusCode), nil
}

// SaveDbStatus saves status code for specified key
func (s *StatusStore) SaveDbStatus(key string, statusCode int) error {
	db, err := leveldb.OpenFile(s.databaseFile, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutVarint(buf, int64(statusCode))

	db.Put([]byte(key), buf, nil)

	return nil
}
