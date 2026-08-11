package httpclient

import (
	"net/http"
	"time"
)

// New creates an HTTP client for inter-service REST calls.
func New(timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	return &http.Client{Timeout: timeout}
}
