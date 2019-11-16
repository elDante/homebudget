package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Split model
type Split struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	TransactionID uuid.UUID `gorm:"type:uuid;not null;"`
	Transaction   Transaction
	AccountID     uuid.UUID `gorm:"type:uuid;not null;"`
	Account       Account
	Value         int
	ValueDenom    int
}

// BeforeCreate will set a UUID rather than numeric ID.
func (split *Split) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
