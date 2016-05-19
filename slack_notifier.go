package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

var slackExpectedKeys = []string{"type", "check_url", "token", "channel"}

type SlackNotifier struct {
	token    string
	userName string
	channel  string
}

func NewSlackNotifier(token string, userName string, channel string) *SlackNotifier {
	s := new(SlackNotifier)
	s.token = token
	s.channel = channel

	if len(userName) == 0 {
		s.userName = "zatsu_monitor"
	} else{
		s.userName = userName
	}

	return s
}

func (s SlackNotifier) ExpectedKeys() []string {
	return slackExpectedKeys
}

func (s SlackNotifier) PostStatus(checkUrl string, beforeStatusCode int, currentStatusCode int) error {
	var statusText, iconEmoji, userName string

	successful := IsSuccessfulStatus(currentStatusCode)

	if successful {
		statusText = "ok"
		iconEmoji = ":green_heart:"
		userName = s.userName + " Successful"
	} else {
		statusText = "down"
		iconEmoji = ":broken_heart:"
		userName = s.userName + " Failure"
	}

	format := `%s is %s
statusCode: %d -> %d`
	message := fmt.Sprintf(format, checkUrl, statusText, beforeStatusCode, currentStatusCode)

	params := slack.NewPostMessageParameters()
	params.Username = userName
	params.IconEmoji = iconEmoji

	api := slack.New(s.token)

	_, _, err := api.PostMessage(s.channel, message, params)

	return err
}
