package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtil_GetStatusCode_Ok(t *testing.T) {
	actual, err := GetStatusCode("https://www.google.co.jp/")

	assert.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestUtil_GetStatusCode_HttpError(t *testing.T) {
	actual, err := GetStatusCode("https://www.google.co.jp/aaa")

	assert.NoError(t, err)
	assert.Equal(t, 404, actual)
}

func TestUtil_GetStatusCode_NoSuchHost(t *testing.T) {
	actual, err := GetStatusCode("https://aaaaaaaaaaaaaaa")

	assert.Error(t, err)
	assert.Equal(t, 0, actual)
}

func TestUtil_IsSuccessful(t *testing.T) {
	assert.Equal(t, false, IsSuccessfulStatus(0))
	assert.Equal(t, true, IsSuccessfulStatus(200))
	assert.Equal(t, true, IsSuccessfulStatus(302))
	assert.Equal(t, false, IsSuccessfulStatus(404))
	assert.Equal(t, false, IsSuccessfulStatus(502))
}
