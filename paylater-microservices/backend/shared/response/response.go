package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// These helpers preserve existing API shapes used across services:
// success/message envelopes and {"error":"..."} for failures.

func JSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}

func Success(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, payload)
}

func Created(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func CreatedPayload(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusCreated, payload)
}

func BadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func Unauthorized(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err})
}

func Forbidden(c *gin.Context, err string) {
	c.JSON(http.StatusForbidden, gin.H{"error": err})
}

func NotFound(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, gin.H{"error": err})
}

func Conflict(c *gin.Context, err string) {
	c.JSON(http.StatusConflict, gin.H{"error": err})
}

func InternalServerError(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func ServiceUnavailable(c *gin.Context, err string) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
}
