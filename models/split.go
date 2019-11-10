package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Split model
type Split struct {
	GUID        uuid.UUID `gorm:"primary_key"`
	Transaction Transaction
	Account     Account
	Value       int
	ValueDenom  int
}

// BeforeCreate set guid before create
func (split *Split) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("GUID", uuid.NewV4().String())
	return nil
}
