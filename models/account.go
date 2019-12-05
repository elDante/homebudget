package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Account model for storing multiple accounts
type Account struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `json:"name"`
	AccountType string    `json:"account_type"`
	CurrencyID  uuid.UUID `gorm:"type:uuid;not null;" json:"currency_id"`
	Currency    Currency  `json:"-"`
	ParentID    uuid.UUID `json:"parent_id"`
	Description string    `json:"description"`
	Placeholder bool      `json:"placeholder"`
}

// AccountPayload proxy structure for create account and initial transaction
type AccountPayload struct {
	Name        string    `json:"name"`
	AccountType string    `json:"account_type"`
	CurrencyID  uuid.UUID `json:"currency_id"`
	ParentID    uuid.UUID `json:"parent_id"`
	Description string    `json:"description"`
	Placeholder bool      `json:"placeholder"`
	Balance     int       `json:"balance"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// CreateAccount create account with started balance
func CreateAccount(db *gorm.DB, data *AccountPayload) error {
	var currency Currency
	if err := db.Where("id = ?", data.CurrencyID).Take(&currency).Error; err != nil {
		return fmt.Errorf("Currency with id: %s does not exists", data.CurrencyID)
	}
	// Create ROOT account if it isn't exists
	var parent Account
	if data.ParentID == uuid.Nil {
		if err := db.Where("account_type = ?", "ROOT").Take(&parent).Error; err != nil {
			parent = Account{Name: "Root Account", AccountType: "ROOT", CurrencyID: currency.ID, ParentID: uuid.Nil, Description: "", Placeholder: false}
			db.Create(&parent)
		}
		data.ParentID = parent.ID
	} else {
		if err := db.Where("id = ?", data.ParentID).Take(&parent).Error; err != nil {
			return fmt.Errorf("Parent account does not exist")
		}
	}

	if err := db.Where("name = ? AND account_type = ? AND parent_id = ?", data.Name, data.AccountType, data.ParentID).Take(&Account{}).Error; err == nil {
		return fmt.Errorf("Account already exists")
	}

	newAccount := Account{Name: data.Name, AccountType: data.AccountType, CurrencyID: currency.ID, ParentID: data.ParentID, Description: data.Description, Placeholder: data.Placeholder}
	db.Create(&newAccount)
	db.Model(&newAccount).Related(&newAccount.Currency)
	if data.Balance != 0 {
		var creditAccount Account
		if err := db.Where("account_type = ? AND currency_id = ?", "EQUITY", data.CurrencyID).First(&creditAccount).Error; err != nil {
			var root Account
			db.Where("account_type = ?", "ROOT").Take(&root)
			creditAccount = Account{Name: fmt.Sprintf("Start balance - %s", currency.Mnemonic), AccountType: "EQUITY", CurrencyID: currency.ID, ParentID: root.ID, Description: "", Placeholder: false}
			db.Create(&creditAccount)
		}
		db.Model(&creditAccount).Related(&creditAccount.Currency)
		CreateTransaction(db, time.Now(), "Start balance", &creditAccount, &newAccount, -data.Balance, data.Balance)
	}
	return nil
}
