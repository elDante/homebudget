package router

import (
	"github.com/jinzhu/gorm"

	"github.com/elDante/homebudget/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Test(c *gin.Context) {
	c.JSON(200, "authorized")
}

// Router create and return gin router
func Router(db *gorm.DB, redis *redis.Client) *gin.Engine {
	r := gin.Default()

	r.POST("/login/", controllers.UserLogin(db, redis))
	v1 := r.Group("/api/v1/")
	v1.Use(NeedAuthorizationMiddleware(redis))
	v1.GET("/", Test)
	return r
}
