package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// NeedAuthorizationMiddleware check authorization cookie
func NeedAuthorizationMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Cookie("user_session")
		c.Next()
	}
}
