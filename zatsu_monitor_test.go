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

func TestZatsuMonitor_HttpStatusCode_Ok(t *testing.T) {
	actual, err := HttpStatusCode("https://www.google.co.jp/")

	assert.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestZatsuMonitor_HttpStatusCodeg_Ng(t *testing.T) {
	actual, err := HttpStatusCode("https://www.google.co.jp/aaa")

	assert.NoError(t, err)
	assert.Equal(t, 404, actual)
}

func DeleteData(key string) {
	db, err := leveldb.OpenFile(TEST_DB_FILE, nil)
	if err != nil {
		panic("Failed: OpenFile " + TEST_DB_FILE)
	}
	defer db.Close()

	db.Delete([]byte(key), nil)
}

func TestZatsuMonitor_GetDbStatus_Exists(t *testing.T) {
	z := NewTestZatsuMonitor()
	z.SaveDbStatus("key", 200)
	defer DeleteData("key")

	actual, err := z.GetDbStatus("key")

	assert.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestZatsuMonitor_GetDbStatus_NotExists(t *testing.T) {
	z := NewTestZatsuMonitor()
	actual, err := z.GetDbStatus("key")

	assert.NoError(t, err)
	assert.Equal(t, 0, actual)
}

func TestZatsuMonitor_IsSuccessful(t *testing.T) {
	assert.Equal(t, false, IsSuccessfulStatus(0))
	assert.Equal(t, true, IsSuccessfulStatus(200))
	assert.Equal(t, true, IsSuccessfulStatus(302))
	assert.Equal(t, false, IsSuccessfulStatus(404))
	assert.Equal(t, false, IsSuccessfulStatus(502))
}
