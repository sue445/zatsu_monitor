package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfigFromData(t *testing.T) {
	yamlData := `name1:
  check_url: "http://example.com/1"
  type: slack
  api_token: "xoxp-0000000000-0000000000-0000000000-000000"
  webhook_url: "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX"
  user_name: "zatsu_monitor"
  channel: "#general"
name2:
  check_url: "http://example.com/2"
  type: chatwork
  api_token: "AAAAAAAA"
  room_id: 111111`

	config, err := LoadConfigFromData(yamlData)

	require.NoError(t, err)

	assert.Equal(t, "http://example.com/1", config["name1"]["check_url"])
	assert.Equal(t, "slack", config["name1"]["type"])
	assert.Equal(t, "xoxp-0000000000-0000000000-0000000000-000000", config["name1"]["api_token"])
	assert.Equal(t, "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX", config["name1"]["webhook_url"])
	assert.Equal(t, "zatsu_monitor", config["name1"]["user_name"])
	assert.Equal(t, "#general", config["name1"]["channel"])

	assert.Equal(t, "http://example.com/2", config["name2"]["check_url"])
	assert.Equal(t, "chatwork", config["name2"]["type"])
	assert.Equal(t, "AAAAAAAA", config["name2"]["api_token"])
	assert.Equal(t, "111111", config["name2"]["room_id"])
}

func TestLoadConfigFromData2(t *testing.T) {
	yamlData := `name1: &common
  check_url: "http://example.com/1"
  type: slack
  api_token: "xoxp-0000000000-0000000000-0000000000-000000"
  webhook_url: "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX"
  user_name: "zatsu_monitor"
  channel: "#general"
name2:
  <<: *common
  channel: "#random"`

	config, err := LoadConfigFromData(yamlData)

	require.NoError(t, err)

	assert.Equal(t, "http://example.com/1", config["name1"]["check_url"])
	assert.Equal(t, "slack", config["name1"]["type"])
	assert.Equal(t, "xoxp-0000000000-0000000000-0000000000-000000", config["name1"]["api_token"])
	assert.Equal(t, "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX", config["name1"]["webhook_url"])
	assert.Equal(t, "zatsu_monitor", config["name1"]["user_name"])
	assert.Equal(t, "#general", config["name1"]["channel"])

	assert.Equal(t, "http://example.com/1", config["name2"]["check_url"])
	assert.Equal(t, "slack", config["name2"]["type"])
	assert.Equal(t, "xoxp-0000000000-0000000000-0000000000-000000", config["name2"]["api_token"])
	assert.Equal(t, "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX", config["name2"]["webhook_url"])
	assert.Equal(t, "zatsu_monitor", config["name2"]["user_name"])
	assert.Equal(t, "#random", config["name2"]["channel"])
}

func TestLoadConfigFromFile(t *testing.T) {
	config, err := LoadConfigFromFile("test/config.yml")

	require.NoError(t, err)

	assert.Equal(t, "http://example.com/1", config["name1"]["check_url"])
	assert.Equal(t, "slack", config["name1"]["type"])
	assert.Equal(t, "xoxp-0000000000-0000000000-0000000000-000000", config["name1"]["api_token"])
	assert.Equal(t, "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXX", config["name1"]["webhook_url"])
	assert.Equal(t, "zatsu_monitor", config["name1"]["user_name"])
	assert.Equal(t, "#general", config["name1"]["channel"])
	assert.Equal(t, "", config["name1"]["check_only_top_of_status_code"])

	assert.Equal(t, "http://example.com/2", config["name2"]["check_url"])
	assert.Equal(t, "chatwork", config["name2"]["type"])
	assert.Equal(t, "AAAAAAAA", config["name2"]["api_token"])
	assert.Equal(t, "111111", config["name2"]["room_id"])

	assert.Equal(t, "true", config["name3"]["check_only_top_of_status_code"])

	assert.Equal(t, "false", config["name4"]["check_only_top_of_status_code"])
}
