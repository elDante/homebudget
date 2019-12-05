package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elDante/homebudget/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/oleiade/reflections.v1"
)

type getResponse struct {
	Transaction models.Transaction `json:"transaction"`
	Credit      models.Split       `json:"credit"`
	Debit       models.Split       `json:"debit"`
}

// GetTransaction return transaction by id
func GetTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var transaction models.Transaction
		var splits []models.Split
		var credit models.Split
		var debit models.Split
		if err := db.Where("id = ?", id).First(&transaction).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			db.Where("transaction_id = ?", transaction.ID).Find(&splits)
			if splits[0].Value < splits[1].Value {
				credit = splits[0]
				debit = splits[1]
			} else {
				credit = splits[1]
				debit = splits[0]
			}
			response := getResponse{Transaction: transaction, Credit: credit, Debit: debit}
			c.JSON(http.StatusOK, &response)
		}
	}
}

type transactionPayload struct {
	PostDate        string `json:"post_date"`
	Description     string `json:"description"`
	CreditAccountID string `json:"credit_account_id"`
	DebitAccountID  string `json:"debit_account_id"`
	CreditValue     int    `json:"credit_value"`
	DebitValue      int    `json:"debit_value"`
}

// CreateTransaction create transaction
func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload transactionPayload
		if err := c.BindJSON(&payload); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			required := []string{"PostDate", "CreditAccountID", "DebitAccountID", "CreditValue", "DebitValue"}
			for _, field := range required {
				if value, _ := reflections.GetField(payload, field); value == nil {
					tag, _ := reflections.GetFieldTag(payload, field, "json")
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Field %s is required", tag)})
					return
				}
			}
			postDateLayout := "2006-01-02"
			postDate, err := time.Parse(postDateLayout, payload.PostDate)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed parse post_date"})
				return
			}
			var creditAccount models.Account
			var debitAccount models.Account
			if err := db.Where("id = ?", payload.CreditAccountID).Take(&creditAccount).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Credit account doesn't exists"})
				return
			}
			db.Model(&creditAccount).Related(&creditAccount.Currency)
			if err := db.Where("id = ?", payload.DebitAccountID).Take(&debitAccount).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Debit account doesn't exists"})
				return
			}
			db.Model(&debitAccount).Related(&debitAccount.Currency)
			if creditAccount.CurrencyID == debitAccount.CurrencyID {
				if (payload.CreditValue + payload.DebitValue) != 0 {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Credit and debit value mismatch"})
					return
				}
			}
			if err := models.CreateTransaction(db, postDate, payload.Description, &creditAccount, &debitAccount, payload.CreditValue, payload.DebitValue); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Transaction succesfuly commited"})
		}
	}
}

// DeleteTransaction delete transaction by id
func DeleteTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var transaction models.Transaction
		if err := db.Where("id = ?", id).First(&transaction).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			db.Where("transaction_id = ?", transaction.ID).Delete(&models.Split{})
			db.Delete(&transaction)
			c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
		}
	}
}
