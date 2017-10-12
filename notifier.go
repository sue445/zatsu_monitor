package main

// PostStatusParam represents request parameter for PostStatus
type PostStatusParam struct {
	CheckUrl                            string
	BeforeStatusCode, CurrentStatusCode int
	HttpError                           error
	ResponseTime                        float64
}

// Notifier represents interface for generic notifier
type Notifier interface {
	// PostStatus perform posting current status for URL
	PostStatus(*PostStatusParam) error

	// ExpectedKeys returns expected keys for SlackNotifier
	ExpectedKeys() []string
}
