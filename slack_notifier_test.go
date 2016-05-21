package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

func NewTestSlackNotifier() *SlackNotifier {
	godotenv.Load()

	apiToken := os.Getenv("SLACK_API_TOKEN")
	userName := os.Getenv("SLACK_USER_NAME")
	channel := os.Getenv("SLACK_CHANNEL")

	if len(userName) == 0 {
		userName = "zatsu_monitor"
	}

	if len(apiToken) == 0 || len(channel) == 0 {
		return nil
	}

	return NewSlackNotifier(apiToken, userName, "#"+channel)
}

func TestSlackNotifier_PostStatus_Successful(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://www.google.co.jp/", 500, 200, nil)
	assert.NoError(t, err)
}

func TestSlackNotifier_PostStatus_Failure(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://www.google.co.jp/aaa", 0, 404, nil)
	assert.NoError(t, err)
}

func TestSlackNotifier_PostStatus_HasError(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://aaaaaaaaa/", 0, 0, errors.New("Test"))
	assert.NoError(t, err)
}
