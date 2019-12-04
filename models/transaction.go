package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Transaction model
type Transaction struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	PostDate    time.Time `json:"post_date"`
	EnterDate   time.Time `json:"enter_date"`
	Description string    `json:"description"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (transaction *Transaction) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// CreateTransaction create transaction and chained splits
func CreateTransaction(db *gorm.DB, postDate time.Time, description string, credit *Account, debit *Account, creditValue int, debitValue int) error {
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err := trx.Error; err != nil {
		trx.Rollback()
		return err
	}

	t := Transaction{PostDate: postDate, EnterDate: time.Now(), Description: description}
	if err := trx.Create(&t).Error; err != nil {
		trx.Rollback()
		return fmt.Errorf("Transaction initialization failed: %s", err.Error())
	}

	if err := trx.Create(&Split{TransactionID: t.ID, AccountID: credit.ID, CurrencyID: credit.CurrencyID, Value: (creditValue * credit.Currency.Fraction), ValueDenom: credit.Currency.Fraction}).Error; err != nil {
		trx.Rollback()
		return fmt.Errorf("Error creating credit transaction: %s", err.Error())
	}

	if err := trx.Create(&Split{TransactionID: t.ID, AccountID: debit.ID, CurrencyID: debit.CurrencyID, Value: (debitValue * debit.Currency.Fraction), ValueDenom: debit.Currency.Fraction}).Error; err != nil {
		trx.Rollback()
		return fmt.Errorf("Error creating debit transaction: %s", err.Error())
	}

	return trx.Commit().Error
}
