package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"paylater.dev/shared/constants"
)

// Middleware ensures every request has an X-Request-ID.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(constants.HeaderRequestID)
		if id == "" {
			id = uuid.NewString()
		}
		c.Set("request_id", id)
		c.Writer.Header().Set(constants.HeaderRequestID, id)
		c.Request.Header.Set(constants.HeaderRequestID, id)
		c.Next()
	}
}

// FromContext returns the request id from Gin context.
func FromContext(c *gin.Context) string {
	if v, ok := c.Get("request_id"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return c.GetHeader(constants.HeaderRequestID)
}
