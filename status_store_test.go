package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

const TEST_DB_FILE = "tmp/zatsu"

func NewTestStatusStore() *StatusStore {
	return NewStatusStore(TEST_DB_FILE)
}

func DeleteData(key string) {
	db, err := leveldb.OpenFile(TEST_DB_FILE, nil)
	if err != nil {
		panic("Failed: OpenFile " + TEST_DB_FILE)
	}
	defer db.Close()

	db.Delete([]byte(key), nil)
}

func TestStatusStore_GetDbStatus_Exists(t *testing.T) {
	store := NewTestStatusStore()
	store.SaveDbStatus("key", 200)
	defer DeleteData("key")

	actual, err := store.GetDbStatus("key")

	assert.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestStatusStore_GetDbStatus_NotExists(t *testing.T) {
	store := NewTestStatusStore()
	actual, err := store.GetDbStatus("key")

	assert.NoError(t, err)
	assert.Equal(t, NotFoundKey, actual)
}
