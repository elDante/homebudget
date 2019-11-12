package database

import (
	"fmt"
	"log"

	"github.com/elDante/homebudget/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

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
		models.CurrencyFixtures(db)
	} else {
		db.AutoMigrate(&models.Currency{})
	}

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

// DBConnector connect to database
func DBConnector(dbinfo *Database) (db *gorm.DB) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8", dbinfo.Username, dbinfo.Password, dbinfo.Host, dbinfo.Port, dbinfo.Name))
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	return
}
