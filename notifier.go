package main

// PostStatusParam represents request parameter for PostStatus
type PostStatusParam struct {
	CheckURL                            string
	BeforeStatusCode, CurrentStatusCode int
	HTTPError                           error
	ResponseTime                        float64
}

// Notifier represents interface for generic notifier
type Notifier interface {
	// PostStatus perform posting current status for URL
	PostStatus(*PostStatusParam) error

	// ExpectedKeys returns expected keys for SlackNotifier
	ExpectedKeys() []string
}
