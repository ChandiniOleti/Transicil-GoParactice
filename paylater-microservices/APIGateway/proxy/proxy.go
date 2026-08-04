package proxy

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// Forward proxies the incoming request to targetBaseURL, preserving
// method, path, query, headers, and body. Downstream response is returned as-is.
func Forward(targetBaseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := strings.TrimRight(targetBaseURL, "/") + c.Request.URL.RequestURI()

		var body io.Reader
		if c.Request.Body != nil {
			body = c.Request.Body
		}

		req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, targetURL, body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "Failed to create upstream request",
			})
			return
		}

		for key, values := range c.Request.Header {
			if isHopByHopHeader(key) {
				continue
			}
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			if isUnavailable(err) {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"error": "Service Unavailable",
				})
				return
			}
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "Upstream request failed",
			})
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			if isHopByHopHeader(key) {
				continue
			}
			for _, value := range values {
				c.Writer.Header().Add(key, value)
			}
		}

		c.Status(resp.StatusCode)
		_, _ = io.Copy(c.Writer, resp.Body)
	}
}

func isHopByHopHeader(header string) bool {
	switch strings.ToLower(header) {
	case "connection", "keep-alive", "proxy-authenticate", "proxy-authorization",
		"te", "trailers", "transfer-encoding", "upgrade", "content-length":
		return true
	default:
		return false
	}
}

func isUnavailable(err error) bool {
	if err == nil {
		return false
	}
	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "connection refused") ||
		strings.Contains(msg, "no such host") ||
		strings.Contains(msg, "dial tcp") ||
		strings.Contains(msg, "connectex")
}
