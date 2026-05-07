package domain

import (
	"errors"
	"time"
)

var (
	ErrTransactionWithoutEntries = errors.New("transaction must have at least one entry")
	ErrTransactionNotBalanced    = errors.New("transaction entries must be balanced")
	ErrInvalidEntryAmount        = errors.New("entry amount must be greater than zero")
)

type Transaction struct {
	ID        string
	Name      string
	Entries   []Entry
	CreatedAt time.Time
}

func NewTransaction(id, name string, entries []Entry, createdAt time.Time) (*Transaction, error) {
	transaction := &Transaction{
		ID:        id,
		Name:      name,
		Entries:   entries,
		CreatedAt: createdAt,
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if len(t.Entries) == 0 {
		return ErrTransactionWithoutEntries
	}

	var debitTotal int64
	var creditTotal int64

	for _, entry := range t.Entries {
		if entry.Amount <= 0 {
			return ErrInvalidEntryAmount
		}

		if !entry.Direction.IsValid() {
			return ErrInvalidDirection
		}

		switch entry.Direction {
		case Debit:
			debitTotal += entry.Amount
		case Credit:
			creditTotal += entry.Amount
		}
	}

	if debitTotal != creditTotal {
		return ErrTransactionNotBalanced
	}

	return nil
}
