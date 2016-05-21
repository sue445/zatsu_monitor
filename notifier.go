package main

type PostStatusParam struct {
	CheckUrl                            string
	BeforeStatusCode, CurrentStatusCode int
	HttpError                           error
}

type Notifier interface {
	PostStatus(PostStatusParam) error
	ExpectedKeys() []string
}
