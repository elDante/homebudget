package router

import (
	"github.com/elDante/homebudget/config"
	"github.com/jinzhu/gorm"

	"github.com/elDante/homebudget/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Router create and return gin router
func Router(db *gorm.DB, redis *redis.Client, site *config.Site) *gin.Engine {
	r := gin.Default()

	r.POST("/login/", controllers.UserLogin(db, redis, site))
	api := r.Group("/api")
	// api.Use(NeedAuthorizationMiddleware(redis))

	currencies := api.Group("/currencies")
	currencies.GET("/", controllers.GetCurrensies(db))
	currencies.POST("/", controllers.CreateCurrency(db))
	currencies.GET("/:id/", controllers.GetCurrency(db))
	currencies.DELETE("/:id/", controllers.DeleteCurrency(db))

	accounts := api.Group("/accounts")
	accounts.GET("/", controllers.GetAccounts(db))
	accounts.POST("/", controllers.CreateAccount(db))
	accounts.GET("/:id/", controllers.GetAccount(db))
	accounts.PUT("/:id/", controllers.UpdateAccount(db))
	accounts.DELETE("/:id/", controllers.DeleteAccount(db))
	accounts.GET("/:id/balance/", controllers.GetAccountBalance(db))
	accounts.GET("/:id/transactions/", controllers.GetAccountTransactions(db))

	transactions := api.Group("/transactions")
	transactions.POST("/", controllers.CreateTransaction(db))
	transactions.GET(":id/", controllers.GetTransaction(db))
	transactions.DELETE("/:id/", controllers.DeleteTransaction(db))

	return r
}
