package main

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var notifier *ChatworkNotifier

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		// If not found .env, skip test
		return
	}

	apiToken := os.Getenv("CHATWORK_API_TOKEN")
	roomId := os.Getenv("CHATWORK_ROOM_ID")
	notifier = NewChatworkNotifier(apiToken, roomId)

	os.Exit(m.Run())
}

func TestChatworkNotifier_PostStatus_True(t *testing.T) {
	notifier.PostStatus("https://www.google.co.jp/", 200, true)
}

func TestChatworkNotifier_PostStatus_False(t *testing.T) {
	notifier.PostStatus("https://www.google.co.jp/aaa", 404, false)
}
