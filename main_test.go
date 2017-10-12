package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type isNotifyFixture struct {
	checkOnlyTopOfStatusCode bool
	beforeStatusCode         int
	currentStatusCode        int
	expected                 bool
}

var isNotifyFixtures = []isNotifyFixture{
	{false, NotFoundKey, 200, false},
	{false, 200, 200, false},
	{false, 200, 500, true},
	{false, 500, 501, true},
	{false, 200, 201, true},
	{true, NotFoundKey, 200, false},
	{true, 200, 200, false},
	{true, 200, 500, true},
	{true, 500, 501, false},
	{true, 200, 201, false},
}

func TestIsNotify(t *testing.T) {
	for _, fixture := range isNotifyFixtures {
		actual := isNotify(fixture.beforeStatusCode, fixture.currentStatusCode, fixture.checkOnlyTopOfStatusCode)
		assert.Equal(t, fixture.expected, actual)
	}
}
