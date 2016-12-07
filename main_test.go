package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type isNotifyFixture struct {
	beforeStatusCode  int
	currentStatusCode int
	expected          bool
}

var isNotifyFixtures = []isNotifyFixture{
	{NOT_FOUND_KEY, 200, false},
	{200, 200, false},
	{200, 500, true},
	{500, 501, true},
}

func TestIsNotify(t *testing.T) {
	for _, fixture := range isNotifyFixtures {
		actual := isNotify(fixture.beforeStatusCode, fixture.currentStatusCode)
		assert.Equal(t, fixture.expected, actual)
	}
}
