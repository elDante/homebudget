package controllers

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/elDante/homebudget/config"

	"github.com/elDante/homebudget/contrib"
	"github.com/elDante/homebudget/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type jsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UserLogin authentificate user
func UserLogin(db *gorm.DB, redis *redis.Client, site *config.Site) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Request.Cookie("authorization"); err != nil {
			c.Request.ParseForm()
			user := models.User{}
			email := c.PostForm("email")
			password := c.PostForm("password")
			db.Where("email = ?", email).First(&user)
			if user.Email == "" {
				c.AbortWithStatusJSON(403, jsonResponse{Code: 403, Message: "Invalid email"})
			}
			if user.Password == contrib.SecretString(password, site.Secret) {
				c.AbortWithStatusJSON(403, jsonResponse{Code: 403, Message: "Invalid password"})
			}
			// All ok - create session
			now, _ := time.Now().MarshalBinary()
			value := sha256.Sum256(now)
			expire := 7776000
			redis.Set(fmt.Sprintf("%x", value), user.ID, (time.Duration(expire) * time.Second))
			c.SetCookie("authorization", fmt.Sprintf("%x", value), expire, "/", "127.0.0.1", false, true)
			c.JSON(200, fmt.Sprintf("Success!, %s", err))
		} else {
			c.JSON(200, jsonResponse{Code: 200, Message: "You are already logged"})
		}
	}
}
