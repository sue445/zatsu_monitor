package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZatsuMonitor_CheckUrl_Ok(t *testing.T) {
	z := new(ZatsuMonitor)
	actual := z.CheckUrl("https://www.google.co.jp/")
	assert.Equal(t, true, actual)
}

func TestZatsuMonitor_CheckUrl_Ng(t *testing.T) {
	z := new(ZatsuMonitor)
	actual := z.CheckUrl("https://www.google.co.jp/aaa")
	assert.Equal(t, false, actual)
}
