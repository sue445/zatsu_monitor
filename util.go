package main

import "net/http"

func GetStatusCode(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func IsSuccessfulStatus(statusCode int) bool {
	n := statusCode / 100

	// Successful: 2xx, 3xx
	return n == 2 || n == 3
}
