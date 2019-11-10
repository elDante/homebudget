package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Account model for storing multiple accounts
type Account struct {
	GUID        uuid.UUID `gorm:"primary_key"`
	Name        string
	AccountType string
	Currency    Currency
	ParentGUID  uuid.UUID
	Description string
	Placeholder bool
}

// BeforeCreate set guid before create
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("GUID", uuid.NewV4().String())
	return nil
}
