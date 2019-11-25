package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Split model
type Split struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TransactionID uuid.UUID `gorm:"type:uuid;not null;" json:"transaction_id"`
	Transaction   Transaction
	AccountID     uuid.UUID `gorm:"type:uuid;not null;" json:"account_id"`
	Account       Account
	Value         int `json:"value"`
	ValueDenom    int `json:"value_denom"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (split *Split) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
