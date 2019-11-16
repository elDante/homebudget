package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Currency model for currencies store
type Currency struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Mnemonic string
	Fullname string
	Fraction int
}

// BeforeCreate will set a UUID rather than numeric ID.
func (currency *Currency) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// CurrencyFixtures creates initial curensies
func CurrencyFixtures(db *gorm.DB) error {
	db.LogMode(true)
	db.Debug().Create(Currency{ID: uuid.NewV4(), Mnemonic: "RUB", Fullname: "Russian rouble", Fraction: 100})
	db.Debug().Create(Currency{ID: uuid.NewV4(), Mnemonic: "EUR", Fullname: "Euro", Fraction: 100})
	db.Debug().Create(Currency{ID: uuid.NewV4(), Mnemonic: "USD", Fullname: "United States dollar", Fraction: 100})
	db.Debug().Create(Currency{ID: uuid.NewV4(), Mnemonic: "GBP", Fullname: "Pound sterling", Fraction: 100})
	return nil
}
