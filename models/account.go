package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Account model for storing multiple accounts
type Account struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `json:"name"`
	AccountType string    `json:"account_type"`
	CurrencyID  uuid.UUID `gorm:"type:uuid;not null;" json:"currency_id"`
	Currency    Currency
	ParentID    uuid.UUID `json:"parent_id"`
	Description string    `json:"description"`
	Placeholder bool      `json:"placeholder"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
