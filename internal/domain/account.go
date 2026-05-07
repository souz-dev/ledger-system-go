package domain

import "time"

type Account struct {
	ID        string
	Name      string
	Direction Direction
	Balance   int64
	CreatedAt time.Time
}

func NewAccount(id, name string, direction Direction, initialBalance int64, createdAt time.Time) *Account {
	return &Account{
		ID:        id,
		Name:      name,
		Direction: direction,
		Balance:   initialBalance,
		CreatedAt: createdAt,
	}
}

func (a *Account) Apply(entry Entry) {
	if a.Direction == entry.Direction {
		a.Balance += entry.Amount
		return
	}

	a.Balance -= entry.Amount
}
