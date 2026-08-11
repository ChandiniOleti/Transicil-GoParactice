package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"paylater.dev/shared/logger"
)

// Run starts an HTTP server with graceful shutdown.
func Run(addr string, handler http.Handler, db *sql.DB) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.L().Info("http server starting", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		logger.L().Error("http server failed", "error", err.Error())
		return err
	case sig := <-stop:
		logger.L().Info("shutdown signal received", "signal", sig.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.L().Error("http shutdown error", "error", err.Error())
	}
	if db != nil {
		if err := db.Close(); err != nil {
			logger.L().Error("database close error", "error", err.Error())
		}
	}
	logger.L().Info("shutdown complete")
	return nil
}

// HealthHandler returns a standard UP health payload.
func HealthHandler(serviceName string, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := "UP"
		httpStatus := http.StatusOK
		if db != nil {
			if err := db.Ping(); err != nil {
				status = "DOWN"
				httpStatus = http.StatusServiceUnavailable
			}
		}
		c.JSON(httpStatus, gin.H{
			"status":  status,
			"service": serviceName,
		})
	}
}
