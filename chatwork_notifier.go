package main

import (
	"fmt"
	chatwork "github.com/griffin-stewie/go-chatwork"
)

var chatworkExpectedKeys = []string{"type", "check_url", "room_id", "api_token"}

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
	return chatworkExpectedKeys
}

func (c ChatworkNotifier) PostStatus(checkUrl string, beforeStatusCode int, currentStatusCode int, httpError error) error {
	chatwork := chatwork.NewClient(c.apiToken)

	var statusText string

	successful := IsSuccessfulStatus(currentStatusCode)

	if successful {
		statusText = "ok (F)"
	} else {
		statusText = "down (devil)"
	}

	title := fmt.Sprintf("%s is %s", checkUrl, statusText)
	body := fmt.Sprintf("statusCode: %d -> %d", beforeStatusCode, currentStatusCode)

	if httpError != nil {
		body += fmt.Sprintf("\nhttpError: %v", httpError)
	}

	message := fmt.Sprintf("[info][title]%s[/title]%s[/info]", title, body)

	_, err := chatwork.PostRoomMessage(c.roomId, message)

	return err
}
