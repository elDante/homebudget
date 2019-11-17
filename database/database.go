package database

import (
	"fmt"
	"log"

	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/models"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// MigrateDB migrate database
func MigrateDB(db *gorm.DB) {
	if !db.HasTable(&models.Account{}) {
		err := db.CreateTable(&models.Account{})
		if err != nil {
			log.Println(err)
		}
	} else {
		db.AutoMigrate(&models.Account{})
	}

	if !db.HasTable(&models.Currency{}) {
		err := db.CreateTable(&models.Currency{})
		if err != nil {
			log.Println(err)
		}
	} else {
		db.AutoMigrate(&models.Currency{})
	}
	models.CurrencyFixtures(db)

	if !db.HasTable(&models.Split{}) {
		err := db.CreateTable(&models.Split{})
		if err != nil {
			log.Println(err)
		}
	} else {
		db.AutoMigrate(&models.Split{})
	}

	if !db.HasTable(&models.Transaction{}) {
		err := db.CreateTable(&models.Transaction{})
		if err != nil {
			log.Println(err)
		}
	} else {
		db.AutoMigrate(&models.Transaction{})
	}
}

// Connector connect to database
func Connector(dbinfo *config.Database) (db *gorm.DB) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbinfo.Host, dbinfo.Port, dbinfo.Username, dbinfo.Name, dbinfo.Password))
	// db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8", dbinfo.Username, dbinfo.Password, dbinfo.Host, dbinfo.Port, dbinfo.Name))
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	return
}
