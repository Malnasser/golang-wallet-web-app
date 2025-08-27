// Package ledger maintains an momery ledger of accounts and transactions
package ledger

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	SAR = "SAR"
	USD = "USD"
)

type Account struct {
	UUID     string
	Name     string
	Balance  int
	Currency string
}

const (
	Debit  = "Debit"
	Credit = "Credit"
	TopUp  = "TopUp"
)

type Transaction struct {
	UUID          string
	AccountID     string
	IdempotencyID string
	TrxType       string
	Amount        int
	AfterBalance  int
	CreatedAt     time.Time
}

type Ledger struct {
	accounts     map[string]*Account
	transactions map[string]*Transaction
}

func NewLedger() *Ledger {
	return &Ledger{
		accounts:     make(map[string]*Account),
		transactions: make(map[string]*Transaction),
	}
}

func (l Ledger) CreateAccount(name string) *Account {
	acc := Account{
		UUID:     uuid.New().String(),
		Name:     name,
		Balance:  0,
		Currency: "SAR",
	}
	l.accounts[acc.UUID] = &acc
	return &acc
}

func (l Ledger) FundTransafer(senderAcc *Account, recieverAcc *Account, amount int) ([2]Transaction, error) {
	if senderAcc.Balance < amount {
		return [2]Transaction{}, errors.New("insufficient funds")
	}

	debitTrx := Transaction{
		UUID:          uuid.New().String(),
		AccountID:     senderAcc.UUID,
		IdempotencyID: uuid.New().String(),
		TrxType:       Debit,
		Amount:        amount,
		AfterBalance:  senderAcc.Balance - amount,
		CreatedAt:     time.Now(),
	}
	senderAcc.Balance = debitTrx.AfterBalance

	creditTrx := Transaction{
		UUID:          uuid.New().String(),
		AccountID:     recieverAcc.UUID,
		IdempotencyID: uuid.New().String(),
		TrxType:       Credit,
		Amount:        amount,
		AfterBalance:  recieverAcc.Balance + amount,
		CreatedAt:     time.Now(),
	}
	recieverAcc.Balance = creditTrx.AfterBalance

	// commit
	l.transactions[debitTrx.UUID] = &debitTrx
	l.transactions[creditTrx.UUID] = &creditTrx

	return [2]Transaction{debitTrx, creditTrx}, nil
}

func (l Ledger) TopUpAccount(acc *Account, amount int) *Transaction {
	topUpTrx := Transaction{
		UUID:          uuid.New().String(),
		AccountID:     acc.UUID,
		IdempotencyID: uuid.New().String(),
		TrxType:       TopUp,
		Amount:        amount,
		AfterBalance:  acc.Balance + amount,
		CreatedAt:     time.Now(),
	}

	acc.Balance = topUpTrx.AfterBalance

	l.transactions[topUpTrx.UUID] = &topUpTrx

	return &topUpTrx
}
