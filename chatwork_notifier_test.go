package main

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func NewTestChatworkNotifier() *ChatworkNotifier {
	godotenv.Load()

	apiToken := os.Getenv("CHATWORK_API_TOKEN")
	roomId := os.Getenv("CHATWORK_ROOM_ID")

	if len(apiToken) == 0 || len(roomId) == 0 {
		return nil
	}
	return NewChatworkNotifier(apiToken, roomId)
}

func TestChatworkNotifier_PostStatus_True(t *testing.T) {
	notifier := NewTestChatworkNotifier()
	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckUrl:          "https://www.google.co.jp/",
		BeforeStatusCode:  500,
		CurrentStatusCode: 200,
		HttpError:         nil,
	}

	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestChatworkNotifier_PostStatus_False(t *testing.T) {
	notifier := NewTestChatworkNotifier()
	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckUrl:          "https://www.google.co.jp/aaa",
		BeforeStatusCode:  0,
		CurrentStatusCode: 404,
		HttpError:         nil,
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}

func TestChatworkNotifier_PostStatus_HasError(t *testing.T) {
	notifier := NewTestChatworkNotifier()
	if notifier == nil {
		return
	}

	param := PostStatusParam{
		CheckUrl:          "https://aaaaaaaaa/",
		BeforeStatusCode:  0,
		CurrentStatusCode: 0,
		HttpError:         errors.New("Test"),
	}
	err := notifier.PostStatus(&param)
	assert.NoError(t, err)
}
