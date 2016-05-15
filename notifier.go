package main

type Notifier interface {
	PostStatus(string, int, bool)
}
