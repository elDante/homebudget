package controllers

import (
	"net/http"

	"github.com/elDante/homebudget/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetCurrensies return all currensies
func GetCurrensies(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var currensies []models.Currency
		db.Find(&currensies)
		c.JSON(http.StatusOK, currensies)
	}
}

// GetCurrency return currency by id
func GetCurrency(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var currency models.Currency
		if err := db.Where("id = ?", id).First(&currency).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			c.JSON(http.StatusOK, currency)
		}
	}
}

// CreateCurrency create currenct
func CreateCurrency(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var currency models.Currency
		if err := c.BindJSON(&currency); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			db.Create(&currency)
			c.JSON(http.StatusCreated, &currency)
		}
	}
}

// DeleteCurrency delete currency by id
func DeleteCurrency(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var currency models.Currency
		if err := db.Where("id = ?", id).First(&currency).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		} else {
			db.Delete(&currency)
			c.JSON(http.StatusOK, gin.H{"message": "Currency deleted"})
		}
	}
}
