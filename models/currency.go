package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Currency model for currencies store
type Currency struct {
	GUID     uuid.UUID `gorm:"primary_key"`
	Mnemonic string
	Fullname string
	Fraction int
}

// BeforeCreate set guid before create
func (currency *Currency) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("GUID", uuid.NewV4().String())
	return nil
}

// CurrencyFixtures creates initial curensies
func CurrencyFixtures(db *gorm.DB) error {
	db.Create(Currency{Mnemonic: "RUB", Fullname: "Russian rouble", Fraction: 100})
	db.Create(Currency{Mnemonic: "EUR", Fullname: "Euro", Fraction: 100})
	db.Create(Currency{Mnemonic: "USD", Fullname: "United States dollar", Fraction: 100})
	db.Create(Currency{Mnemonic: "GBP", Fullname: "Pound sterling", Fraction: 100})
	return nil
}
