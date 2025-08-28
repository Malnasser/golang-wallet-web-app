// Package ledger maintains an momery ledger of accounts and transactions
package ledger

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	UUID        uuid.UUID    `json:"uuid" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AccountName string       `json:"account_name" gorm:"column:account_name;not null;size:255" validate:"required,min=3,max=255"`
	Balance     int64        `json:"balance" gorm:"not null;default:0"`
	Currency    CurrencyType `json:"currency" gorm:"type:currency_type;not null" validate:"required,oneof=SAR USD"`
	CreatedAt   time.Time    `json:"created_at" gorm:"not null;default:current_timestamp"`

	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:AccountUUID;references:UUID"`
}

func (Account) TableName() string {
	return "accounts"
}

func (account *Account) BeforeCreate(tx *gorm.DB) error {
	if account.UUID == uuid.Nil {
		account.UUID = uuid.New()
	}
	return nil
}

func (account *Account) BeforeUpdate(tx *gorm.DB) (err error) {
	if account.Balance < 0 {
		return errors.New("cannot perform update: insuffivient funds")
	}
	return nil
}

type Transaction struct {
	UUID          uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AccountUUID   uuid.UUID       `gorm:"type:uuid;not null"`
	IdempotencyID string          `gorm:"type:varchar(36);not null;uniqueIndex"`
	TrxType       TransactionType `gorm:"type:transaction_type;not null"`
	Amount        int64           `gorm:"not null"`
	AfterBalance  int64           `gorm:"not null"`
	CreatedAt     time.Time       `gorm:"not null;default:current_timestamp"`

	Account Account `gorm:"foreignKey:AccountUUID;references:UUID"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) error {
	if transaction.UUID == uuid.Nil {
		transaction.UUID = uuid.New()
	}
	return nil
}
