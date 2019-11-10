package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Transaction model
type Transaction struct {
	GUID        uuid.UUID `gorm:"primary_key"`
	Currency    Currency
	PostDate    time.Time
	EnterDate   time.Time
	Description string
}

// BeforeCreate set guid before create
func (transaction *Transaction) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("GUID", uuid.NewV4().String())
	return nil
}
