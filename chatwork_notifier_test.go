package main

import (
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

	err := notifier.PostStatus("https://www.google.co.jp/", 0, 200)
	assert.NoError(t, err)
}

func TestChatworkNotifier_PostStatus_False(t *testing.T) {
	notifier := NewTestChatworkNotifier()
	if notifier == nil {
		return
	}

	err := notifier.PostStatus("https://www.google.co.jp/aaa", 0, 404)
	assert.NoError(t, err)
}
