package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Transaction model
type Transaction struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CurrencyID  uuid.UUID `gorm:"type:uuid;not null;"`
	Currency    Currency
	PostDate    time.Time
	EnterDate   time.Time
	Description string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (transaction *Transaction) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
