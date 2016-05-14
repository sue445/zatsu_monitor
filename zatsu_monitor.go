package main

import (
	"net/http"
)

var OK_STATUS_CODES = []int{200, 301, 302, 303, 304, 307, 308}

type ZatsuMonitor struct {
}

func (z ZatsuMonitor) CheckUrl(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		return false
	}

	for _, v := range OK_STATUS_CODES {
		if v == resp.StatusCode {
			return true
		}
	}

	return false
}
