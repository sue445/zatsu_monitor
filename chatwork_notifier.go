package main

import (
	"fmt"
	chatwork "github.com/griffin-stewie/go-chatwork"
)

type ChatworkNotifier struct {
	apiToken string
	roomId   string
}

func NewChatworkNotifier(apiToken string, roomId string) *ChatworkNotifier {
	if len(apiToken) == 0 {
		panic("apiToken is required")
	}

	if len(roomId) == 0 {
		panic("roomId is required")
	}

	c := new(ChatworkNotifier)
	c.apiToken = apiToken
	c.roomId = roomId
	return c
}

func (c ChatworkNotifier) PostStatus(checkUrl string, statusCode int) {
	chatwork := chatwork.NewClient(c.apiToken)

	var statusText string

	successful := IsSuccessfulStatus(statusCode)

	if successful {
		statusText = "ok (F)"
	} else {
		statusText = "down (devil)"
	}

	message := fmt.Sprintf("[info][title]%s is %s[/title]statusCode=%d[/info]", checkUrl, statusText, statusCode)

	_, err := chatwork.PostRoomMessage(c.roomId, message)
	if err != nil {
		panic(fmt.Sprintf("Can not post message: %v", err))
	}
}
