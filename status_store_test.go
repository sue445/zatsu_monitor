package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

const TestDbFile = "tmp/zatsu"

func NewTestStatusStore() *StatusStore {
	return NewStatusStore(TestDbFile)
}

func DeleteData(key string) {
	db, err := leveldb.OpenFile(TestDbFile, nil)
	if err != nil {
		panic("Failed: OpenFile " + TestDbFile)
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
