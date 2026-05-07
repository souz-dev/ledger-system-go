package domain

import "time"

type Entry struct {
	ID        string
	AccountID string
	Direction Direction
	Amount    int64
	CreatedAt time.Time
}

func NewEntry(id, accountID string, direction Direction, amount int64, createdAt time.Time) Entry {
	return Entry{
		ID:        id,
		AccountID: accountID,
		Direction: direction,
		Amount:    amount,
		CreatedAt: createdAt,
	}
}
