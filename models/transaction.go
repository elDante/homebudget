package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Transaction model
type Transaction struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CurrencyID  uuid.UUID `gorm:"type:uuid;not null;" json:"currency_id"`
	Currency    Currency
	PostDate    time.Time `json:"post_date"`
	EnterDate   time.Time `json:"enter_date"`
	Description string    `json:"description"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (transaction *Transaction) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
