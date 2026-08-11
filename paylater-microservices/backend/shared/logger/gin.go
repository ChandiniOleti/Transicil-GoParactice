package logger

import (
	"time"

	"github.com/gin-gonic/gin"

	"paylater.dev/shared/requestid"
)

// GinMiddleware logs method, path, status, duration, and request id.
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		rid := requestid.FromContext(c)
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		WithRequest(rid, c.Request.Method, path).Info("request completed",
			"status", c.Writer.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		)
	}
}
