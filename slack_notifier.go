package main

import (
	"errors"
	"fmt"

	"github.com/slack-go/slack"
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
		api := slack.New(s.apiToken)

		_, _, err := api.PostMessage(
			s.channel,
			slack.MsgOptionUsername(userName),
			slack.MsgOptionIconEmoji(iconEmoji),
			slack.MsgOptionAttachments(slack.Attachment{
				Fallback: message,
				Text:     message,
				Color:    attachmentColor,
			}),
		)
		return err
	} else if len(s.webhookURL) > 0 {
		return slack.PostWebhook(s.webhookURL, &slack.WebhookMessage{
			Channel:   s.channel,
			Username:  userName,
			IconEmoji: iconEmoji,
			Attachments: []slack.Attachment{
				{Fallback: message, Text: message, Color: attachmentColor},
			},
		})

	} else {
		return errors.New("Either `api_token` or `webhook_url` is required")
	}
}
