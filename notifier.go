package main

type PostStatusParam struct {
	CheckUrl                            string
	BeforeStatusCode, CurrentStatusCode int
	HttpError                           error
	ResponseTime                        float64
}

type Notifier interface {
	PostStatus(*PostStatusParam) error
	ExpectedKeys() []string
}
