package main

type Notifier interface {
	PostStatus(string, int, int) error
	ExpectedKeys() []string
}
