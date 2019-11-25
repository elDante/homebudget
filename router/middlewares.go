package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// NeedAuthorizationMiddleware check authorization cookie
func NeedAuthorizationMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized"})
		} else {
			_, err := redis.Get(cookie.Value).Result()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized"})
			}
			c.Next()
		}
	}
}
