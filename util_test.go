package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtil_HttpStatusCode_Ok(t *testing.T) {
	actual, err := HttpStatusCode("https://www.google.co.jp/")

	assert.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestUtil_HttpStatusCodeg_Ng(t *testing.T) {
	actual, err := HttpStatusCode("https://www.google.co.jp/aaa")

	assert.NoError(t, err)
	assert.Equal(t, 404, actual)
}

func TestUtil_IsSuccessful(t *testing.T) {
	assert.Equal(t, false, IsSuccessfulStatus(0))
	assert.Equal(t, true, IsSuccessfulStatus(200))
	assert.Equal(t, true, IsSuccessfulStatus(302))
	assert.Equal(t, false, IsSuccessfulStatus(404))
	assert.Equal(t, false, IsSuccessfulStatus(502))
}
