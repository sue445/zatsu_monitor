package main

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"os"
	"testing"
)

type SlackNotifierMode int

const (
	WithAPIToken = iota
	WithWebhookURL
	Nothing
)

func NewTestSlackNotifier(mode SlackNotifierMode) *SlackNotifier {
	godotenv.Load()

	apiToken := os.Getenv("SLACK_API_TOKEN")
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	userName := os.Getenv("SLACK_USER_NAME")
	channel := os.Getenv("SLACK_CHANNEL")

	if len(userName) == 0 {
		userName = "zatsu_monitor"
	}

	if len(apiToken) == 0 || len(webhookURL) == 0 || len(channel) == 0 {
		return nil
	}

	switch mode {
	case WithAPIToken:
		return NewSlackNotifier(apiToken, "", userName, "#"+channel)
	case WithWebhookURL:
		return NewSlackNotifier("", webhookURL, userName, "#"+channel)
	case Nothing:
		return NewSlackNotifier("", "", userName, "#"+channel)
	default:
		return NewSlackNotifier("", "", userName, "#"+channel)
	}
}

func TestSlackNotifier_WithApiToken_PostStatus_Successful(t *testing.T) {
	notifier := NewTestSlackNotifier(WithAPIToken)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.co.jp/",
		BeforeStatusCode:  500,
		CurrentStatusCode: 200,
		HTTPError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_WithApiToken_PostStatus_Failure(t *testing.T) {
	notifier := NewTestSlackNotifier(WithAPIToken)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.co.jp/aaa",
		BeforeStatusCode:  0,
		CurrentStatusCode: 404,
		HTTPError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_WithApiToken_PostStatus_HasError(t *testing.T) {
	notifier := NewTestSlackNotifier(WithAPIToken)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://aaaaaaaaa/",
		BeforeStatusCode:  0,
		CurrentStatusCode: 0,
		HTTPError:         errors.New("Test"),
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_WithWebhookURL_PostStatus_Successful(t *testing.T) {
	notifier := NewTestSlackNotifier(WithWebhookURL)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.com/",
		BeforeStatusCode:  500,
		CurrentStatusCode: 200,
		HTTPError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_WithWebhookURL_PostStatus_Failure(t *testing.T) {
	notifier := NewTestSlackNotifier(WithWebhookURL)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.com/aaa",
		BeforeStatusCode:  0,
		CurrentStatusCode: 404,
		HTTPError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_WithWebhookURL_PostStatus_HasError(t *testing.T) {
	notifier := NewTestSlackNotifier(WithWebhookURL)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://bbbbbbbbb/",
		BeforeStatusCode:  0,
		CurrentStatusCode: 0,
		HTTPError:         errors.New("Test"),
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestSlackNotifier_Nothing_PostStatus_Successful(t *testing.T) {
	notifier := NewTestSlackNotifier(Nothing)

	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckURL:          "https://www.google.co.jp/",
		BeforeStatusCode:  500,
		CurrentStatusCode: 200,
		HTTPError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.Error(t, err, "Either `api_token` or `webhook_url` is required")
}
