package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigFromData(t *testing.T) {
	yamlData := `name1:
  check_url: "http://example.com/1"
  type: slack
  webhook_url: "https://hooks.slack.com/services/AAAAAAAAA/BBBBBBBB/CCCCCCCCCCCCC"
  user_name: "zatsu_monitor"
  channel: "#general"
name2:
  check_url: "http://example.com/2"
  type: chatwork
  api_token: "AAAAAAAA"
  room_id: 111111
`

	config, err := LoadConfigFromData(yamlData)

	assert.NoError(t, err)

	assert.Equal(t, "http://example.com/1", config["name1"]["check_url"])
	assert.Equal(t, "slack", config["name1"]["type"])
	assert.Equal(t, "https://hooks.slack.com/services/AAAAAAAAA/BBBBBBBB/CCCCCCCCCCCCC", config["name1"]["webhook_url"])
	assert.Equal(t, "zatsu_monitor", config["name1"]["user_name"])
	assert.Equal(t, "#general", config["name1"]["channel"])

	assert.Equal(t, "http://example.com/2", config["name2"]["check_url"])
	assert.Equal(t, "chatwork", config["name2"]["type"])
	assert.Equal(t, "AAAAAAAA", config["name2"]["api_token"])
	assert.Equal(t, "111111", config["name2"]["room_id"])
}
