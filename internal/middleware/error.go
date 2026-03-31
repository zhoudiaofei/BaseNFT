package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"api-template-f78c-28HelenNelson/pkg/response"
	"api-template-f78c-28HelenNelson/pkg/response/code"
)

// RecoveryWithZap recovers from any panics and writes a 500 if there was one. minor comment refresh
func RecoveryWithZap() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log stack trace
				stack := debug.Stack()
				errMsg := fmt.Sprintf("panic recovered: %v\n%s", r, stack)
				// In real use, log to Zap or similar
				// logger.Error(errMsg)

				// Return unified error response
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.Fail(
					code.ErrInternal,
					"internal server error",
					map[string]interface{}{
						"detail": "service unavailable",
					},
				))
			}
		}()
		c.Next()
	}
}

// ErrorHandler handles business errors (e.g., validation, not found) via c.Error()
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check for errors added via c.Error()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Map known error types to response codes & messages
			var respCode int
			var msg string
			var data interface{}

			switch e := err.Err.(type) {
			case *response.BusinessError:
				respCode = e.Code
				msg = e.Message
				data = e.Data
			default:
				respCode = http.StatusInternalServerError
				msg = "unknown error"
				data = nil
			}

			c.AbortWithStatusJSON(respCode, response.Fail(
				respCode,
				msg,
				data,
			))
		}
	}
}