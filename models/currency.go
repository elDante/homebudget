package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Currency model for currencies store
type Currency struct {
	GUID     uuid.UUID `gorm:"primary_key"`
	Mnemonic string
	Fullname string
	Fraction int
}

// BeforeCreate set guid before create
func (currency *Currency) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("GUID", uuid.NewV4().String())
	return nil
}
