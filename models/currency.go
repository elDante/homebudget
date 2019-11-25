package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Currency model for currencies store
type Currency struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Mnemonic string    `json:"mnemonic"`
	Fullname string    `json:"full_name"`
	Fraction int       `json:"fraction"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (currency *Currency) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// CurrencyFixtures creates initial curensies
func CurrencyFixtures(db *gorm.DB) error {
	fixtures := [4]Currency{
		Currency{Mnemonic: "RUB", Fullname: "Russian rouble", Fraction: 100},
		Currency{Mnemonic: "EUR", Fullname: "Euro", Fraction: 100},
		Currency{Mnemonic: "USD", Fullname: "United States dollar", Fraction: 100},
		Currency{Mnemonic: "GBP", Fullname: "Pound sterling", Fraction: 100},
	}
	for _, fixture := range fixtures {
		db.Where(&fixture).FirstOrCreate(&fixture)
	}

	return nil
}
