package controllers

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/elDante/homebudget/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAccounts return all accounts
func GetAccounts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var accounts []models.Account
		db.Find(&accounts)
		c.JSON(http.StatusOK, &accounts)
	}
}

// GetAccount return account by id
func GetAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account
		if err := db.Where("id = ?", id).First(&account).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			c.JSON(http.StatusOK, &account)
		}
	}
}

// CreateAccount create account
func CreateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var metaAccount models.AccountPayload
		if err := c.BindJSON(&metaAccount); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			if err := models.CreateAccount(db, &metaAccount); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, &metaAccount)
			}
		}
	}
}

// UpdateAccount update account by id
func UpdateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account

		if err := db.Where("id = ?", id).First(&account).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			c.BindJSON(&account)
			db.Save(&account)
			c.JSON(http.StatusOK, &account)
		}
	}
}

// DeleteAccount delete account by id
func DeleteAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account
		if err := db.Where("id = ?", id).Take(&account).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			var transactions []models.Split
			if err := db.Where("account_id = ?", id).Find(&transactions).Error; err != nil {
				db.Delete(&account)
				c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
			}
			if len(transactions) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot delete account, account has %d transactions", len(transactions))})
			} else {
				db.Delete(&account)
				c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
			}
		}
	}
}

type accountBalance struct {
	Value      int `json:"value"`
	ValueDenom int `json:"value_denom"`
}

// GetAccountBalance return account balance by id
func GetAccountBalance(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account
		var balance accountBalance
		if err := db.Where("id = ?", id).First(&account).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			db.Model(&models.Split{}).Select("SUM(value) as value, value_denom").Where("account_id = ?", account.ID).Group("value_denom").Scan(&balance)
			c.JSON(http.StatusOK, &balance)
		}
	}
}

type accountTransaction struct {
	ID          uuid.UUID `json:"id"`
	PostDate    time.Time `json:"post_date"`
	EnterDate   time.Time `json:"enter_date"`
	Description string    `json:"description"`
	CurrencyID  uuid.UUID `json:"currency_id"`
	Value       int       `json:"value"`
	ValueDenom  int       `json:"value_denom"`
}

// GetAccountTransactions return all transaction by account id
func GetAccountTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account
		var transaction []accountTransaction
		if err := db.Where("id = ?", id).First(&account).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			db.Table("transactions as t").Select("t.id, t.post_date, t.enter_date, t.description, s.currency_id, s.value, s.value_denom").Joins("JOIN splits AS s ON t.id = s.transaction_id AND s.account_id = ?", account.ID).Scan(&transaction)
			c.JSON(http.StatusOK, &transaction)
		}
	}
}
