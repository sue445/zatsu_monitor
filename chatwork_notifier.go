package main

import (
	"fmt"
	chatwork "github.com/griffin-stewie/go-chatwork"
)

var expectedKeys = []string{"type", "check_url", "room_id", "api_token"}

type ChatworkNotifier struct {
	apiToken string
	roomId   string
}

func NewChatworkNotifier(apiToken string, roomId string) *ChatworkNotifier {
	c := new(ChatworkNotifier)
	c.apiToken = apiToken
	c.roomId = roomId
	return c
}

func (c ChatworkNotifier) ExpectedKeys() []string {
	return expectedKeys
}

func (c ChatworkNotifier) PostStatus(checkUrl string, beforeStatusCode int, currentStatusCode int) error {
	chatwork := chatwork.NewClient(c.apiToken)

	var statusText string

	successful := IsSuccessfulStatus(currentStatusCode)

	if successful {
		statusText = "ok (F)"
	} else {
		statusText = "down (devil)"
	}

	message := fmt.Sprintf("[info][title]%s is %s[/title]statusCode: %d -> %d[/info]", checkUrl, statusText, beforeStatusCode, currentStatusCode)

	_, err := chatwork.PostRoomMessage(c.roomId, message)

	return err
}
