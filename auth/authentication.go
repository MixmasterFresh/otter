package auth

import (
	"github.com/gin-gonic/gin"
)

var masterKey string

// Initialize collects and sets all initial variables
func Initialize(key string) {
	masterKey = key
	StartRandomGeneration()
	StartKeyGeneration(32)
}

// ServerAuthentication authenticates the server via the X-Otter-Key
func ServerAuthentication() gin.HandlerFunc { //gin middleware
	return func(c *gin.Context) {
		key := c.Request.Header.Get("X-Otter-Key")

		if key == "" {
			Unauthorized(c)
			return
		}

		if key != masterKey {
			Forbidden(c)
			return
		}

		c.Next()
	}
}

// Unauthorized implements a generic 401 Unauthorized response
func Unauthorized(c *gin.Context) {
	RespondWithError(401, "You are missing required credentials for this resource", c)
}

// Forbidden implements a generic 403 Forbidden response
func Forbidden(c *gin.Context) {
	RespondWithError(403, "Forbidden", c)
}

// NotFound implements a generic 404 Not Found response
func NotFound(c *gin.Context) {
	RespondWithError(404, "Not Found", c)
}

func Conflict(c *gin.Context, msg string) {
	RespondWithError(409, "Conflict: " + msg, c)
}

// RespondWithError is for all other error responses
func RespondWithError(status int, message string, c *gin.Context) {
	c.String(status, message)
}
