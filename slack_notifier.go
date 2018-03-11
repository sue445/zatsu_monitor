package main

import (
	"errors"
	"fmt"
	"github.com/bluele/slack"
)

var slackExpectedKeys = []string{"type", "check_url", "channel"}

// SlackNotifier represents notifier for Slack
type SlackNotifier struct {
	apiToken   string
	webhookURL string
	userName   string
	channel    string
}

// NewSlackNotifier create new SlackNotifier instance
func NewSlackNotifier(apiToken string, webhookURL string, userName string, channel string) *SlackNotifier {
	s := new(SlackNotifier)
	s.apiToken = apiToken
	s.webhookURL = webhookURL
	s.channel = channel

	if len(userName) == 0 {
		s.userName = "zatsu_monitor"
	} else {
		s.userName = userName
	}

	return s
}

// ExpectedKeys returns expected keys for SlackNotifier
func (s *SlackNotifier) ExpectedKeys() []string {
	return slackExpectedKeys
}

// PostStatus perform posting current status for URL
func (s *SlackNotifier) PostStatus(param *PostStatusParam) error {
	var statusText, iconEmoji, userName, attachmentColor string

	successful := IsSuccessfulStatus(param.CurrentStatusCode)

	if successful {
		statusText = "ok"
		iconEmoji = ":green_heart:"
		userName = s.userName + " Successful"
		attachmentColor = "good"
	} else {
		statusText = "down"
		iconEmoji = ":broken_heart:"
		userName = s.userName + " Failure"
		attachmentColor = "danger"
	}

	format := `%s is %s
statusCode: %d -> %d
responseTime: %f sec`
	message := fmt.Sprintf(format, param.CheckURL, statusText, param.BeforeStatusCode, param.CurrentStatusCode, param.ResponseTime)

	if param.HTTPError != nil {
		message += fmt.Sprintf("\nhttpError: %v", param.HTTPError)
	}

	if len(s.apiToken) > 0 {
		params := slack.ChatPostMessageOpt{
			Username:  userName,
			IconEmoji: iconEmoji,
			Attachments: []*slack.Attachment{
				{Fallback: message, Text: message, Color: attachmentColor},
			},
		}

		api := slack.New(s.apiToken)

		return api.ChatPostMessage(s.channel, "", &params)

	} else if len(s.webhookURL) > 0 {
		hook := slack.NewWebHook(s.webhookURL)
		params := slack.WebHookPostPayload{
			Channel:   s.channel,
			Username:  userName,
			IconEmoji: iconEmoji,
			Attachments: []*slack.Attachment{
				{Fallback: message, Text: message, Color: attachmentColor},
			},
		}
		return hook.PostMessage(&params)

	} else {
		return errors.New("Either `api_token` or `webhook_url` is required")
	}
}
