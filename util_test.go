package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUtil_GetStatusCode_Ok(t *testing.T) {
	actual, err := GetStatusCode("https://www.google.co.jp/")

	require.NoError(t, err)
	assert.Equal(t, 200, actual)
}

func TestUtil_GetStatusCode_HttpError(t *testing.T) {
	actual, err := GetStatusCode("https://www.google.co.jp/aaa")

	require.NoError(t, err)
	assert.Equal(t, 404, actual)
}

func TestUtil_GetStatusCode_NoSuchHost(t *testing.T) {
	actual, err := GetStatusCode("https://aaaaaaaaaaaaaaa")

	require.Error(t, err)
	assert.Equal(t, 0, actual)
}

func TestUtil_IsSuccessful(t *testing.T) {
	assert.False(t, IsSuccessfulStatus(0))
	assert.True(t, IsSuccessfulStatus(200))
	assert.True(t, IsSuccessfulStatus(302))
	assert.False(t, IsSuccessfulStatus(404))
	assert.False(t, IsSuccessfulStatus(502))
}
