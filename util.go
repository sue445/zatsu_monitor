package main

import "net/http"

// GetStatusCode checks and returns status code for URL
func GetStatusCode(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// IsSuccessfulStatus returns whether status code is successful. (2xx or 3xx)
func IsSuccessfulStatus(statusCode int) bool {
	n := statusCode / 100

	// Successful: 2xx, 3xx
	return n == 2 || n == 3
}
