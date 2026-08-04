package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"paylater.dev/shared/response"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Middleware applies a lightweight per-IP token bucket.
// intended for public endpoints only (login/register).
func Middleware(r rate.Limit, burst int) gin.HandlerFunc {
	var (
		mu       sync.Mutex
		visitors = make(map[string]*visitor)
	)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, v := range visitors {
				if time.Since(v.lastSeen) > 3*time.Minute {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		v, ok := visitors[ip]
		if !ok {
			v = &visitor{limiter: rate.NewLimiter(r, burst)}
			visitors[ip] = v
		}
		v.lastSeen = time.Now()
		limiter := v.limiter
		mu.Unlock()

		if !limiter.Allow() {
			response.JSON(c, http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
