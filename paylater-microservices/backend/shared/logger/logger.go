package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	mu     sync.Mutex
	base   *slog.Logger
	service string
)

// Init configures the process-wide structured logger.
func Init(serviceName string) {
	mu.Lock()
	defer mu.Unlock()
	service = serviceName
	base = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", serviceName)
	slog.SetDefault(base)
}

// L returns the service logger.
func L() *slog.Logger {
	mu.Lock()
	defer mu.Unlock()
	if base == nil {
		Init("unknown")
	}
	return base
}

// WithRequest returns a logger enriched with request metadata.
func WithRequest(requestID, method, path string) *slog.Logger {
	return L().With(
		"request_id", requestID,
		"method", method,
		"path", path,
	)
}
