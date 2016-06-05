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

func (c *ChatworkNotifier) ExpectedKeys() []string {
	return chatworkExpectedKeys
}

func (c *ChatworkNotifier) PostStatus(param *PostStatusParam) error {
	chatwork := chatwork.NewClient(c.apiToken)

	var statusText string

	successful := IsSuccessfulStatus(param.CurrentStatusCode)

	if successful {
		statusText = "ok (F)"
	} else {
		statusText = "down (devil)"
	}

	title := fmt.Sprintf("%s is %s", param.CheckUrl, statusText)
	body := fmt.Sprintf("statusCode: %d -> %d", param.BeforeStatusCode, param.CurrentStatusCode)

	if param.HttpError != nil {
		body += fmt.Sprintf("\nhttpError: %v", param.HttpError)
	}

	message := fmt.Sprintf("[info][title]%s[/title]%s[/info]", title, body)

	_, err := chatwork.PostRoomMessage(c.roomId, message)

	return err
}
