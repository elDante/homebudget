package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Account model for storing multiple accounts
type Account struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Name        string
	AccountType string
	CurrencyID  uuid.UUID `gorm:"type:uuid;not null;"`
	Currency    Currency
	ParentID    uuid.UUID
	Description string
	Placeholder bool
}

// BeforeCreate will set a UUID rather than numeric ID.
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
