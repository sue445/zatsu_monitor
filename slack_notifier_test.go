package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func NewTestSlackNotifier() *SlackNotifier {
	godotenv.Load()

	token := os.Getenv("SLACK_TOKEN")
	userName := os.Getenv("SLACK_USER_NAME")
	channel := os.Getenv("SLACK_CHANNEL")

	if len(userName) == 0 {
		userName = "zatsu_monitor"
	}

	if len(token) == 0 || len(channel) == 0 {
		return nil
	}

	return NewSlackNotifier(token, userName, "#"+channel)
}

func TestSlackNotifier_PostStatus_Successful(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://www.google.co.jp/", 0, 200)
	assert.NoError(t, err)
}

func TestSlackNotifier_PostStatus_Failure(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://www.google.co.jp/aaa", 0, 404)
	assert.NoError(t, err)
}
