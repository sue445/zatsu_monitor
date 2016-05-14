package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

const TEST_DB_FILE = "tmp/zatsu"

func NewTestZatsuMonitor() *ZatsuMonitor {
	return NewZatsuMonitor(TEST_DB_FILE)
}

func TestZatsuMonitor_CheckUrl_Ok(t *testing.T) {
	z := NewTestZatsuMonitor()
	actual := z.CheckUrl("https://www.google.co.jp/")
	assert.Equal(t, true, actual)
}

func TestZatsuMonitor_CheckUrl_Ng(t *testing.T) {
	z := NewTestZatsuMonitor()
	actual := z.CheckUrl("https://www.google.co.jp/aaa")
	assert.Equal(t, false, actual)
}

func PutData(key string, value string) {
	db, err := leveldb.OpenFile(TEST_DB_FILE, nil)
	if err != nil {
		panic("Failed: OpenFile " + TEST_DB_FILE)
	}
	defer db.Close()

	db.Put([]byte(key), []byte(value), nil)
}

func DeleteData(key string) {
	db, err := leveldb.OpenFile(TEST_DB_FILE, nil)
	if err != nil {
		panic("Failed: OpenFile " + TEST_DB_FILE)
	}
	defer db.Close()

	db.Delete([]byte(key), nil)
}

func TestZatsuMonitor_GetDbStatus_ExistsTrue(t *testing.T) {
	PutData("key", STATUS_TRUE)
	defer DeleteData("key")

	z := NewTestZatsuMonitor()
	actual := z.GetDbStatus("key")
	assert.Equal(t, true, actual)
}

func TestZatsuMonitor_GetDbStatus_ExistsFalse(t *testing.T) {
	PutData("key", STATUS_FALSE)
	defer DeleteData("key")

	z := NewTestZatsuMonitor()
	actual := z.GetDbStatus("key")
	assert.Equal(t, false, actual)
}

func TestZatsuMonitor_GetDbStatus_NotExists(t *testing.T) {
	z := NewTestZatsuMonitor()
	actual := z.GetDbStatus("key")
	assert.Equal(t, true, actual)
}
