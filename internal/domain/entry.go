package domain

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID        string
	AccountID string
	Direction Direction
	Amount    int64 // Amount represents cents.
	CreatedAt time.Time
}

func NewEntry(id, accountID string, direction Direction, amount int64, createdAt time.Time) (Entry, error) {
	if !direction.IsValid() {
		return Entry{}, ErrInvalidDirection
	}

	if amount <= 0 {
		return Entry{}, ErrInvalidEntryAmount
	}

	if id == "" {
		id = uuid.NewString()
	}

	return Entry{
		ID:        id,
		AccountID: accountID,
		Direction: direction,
		Amount:    amount,
		CreatedAt: createdAt,
	}, nil
}
