package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NeedAuthorizationMiddleware check authorization cookie
func NeedAuthorizationMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := errorResponse{Code: 401, Message: "You are not authorized"}
		cookie, err := c.Request.Cookie("authorization")
		if err != nil {
			c.AbortWithStatusJSON(401, &e)
		} else {
			_, err := redis.Get(cookie.Value).Result()
			if err != nil {
				c.AbortWithStatusJSON(401, &e)
			}
			c.Next()
		}
	}
}
