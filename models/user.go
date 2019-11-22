package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model for storing homebudget user
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Password string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
