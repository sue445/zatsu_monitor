package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"os"
	"testing"
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

	param := PostStatusParam{
		CheckURL:          "https://www.google.co.jp/",
		BeforeStatusCode:  500,
		CurrentStatusCode: 200,
		HttpError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_PostStatus_Failure(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.co.jp/aaa",
		BeforeStatusCode:  0,
		CurrentStatusCode: 404,
		HttpError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_PostStatus_HasError(t *testing.T) {
	notifier := NewTestSlackNotifier()

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://aaaaaaaaa/",
		BeforeStatusCode:  0,
		CurrentStatusCode: 0,
		HttpError:         errors.New("Test"),
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}
