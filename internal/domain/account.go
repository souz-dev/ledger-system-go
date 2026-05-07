package domain

import "time"

type Account struct {
	ID        string
	Name      string
	Direction Direction
	Balance   int64 // Balance represents cents.
	CreatedAt time.Time
}

func NewAccount(id, name string, direction Direction, initialBalance int64, createdAt time.Time) (*Account, error) {
	if !direction.IsValid() {
		return nil, ErrInvalidDirection
	}

	return &Account{
		ID:        id,
		Name:      name,
		Direction: direction,
		Balance:   initialBalance,
		CreatedAt: createdAt,
	}, nil
}

func (a *Account) Apply(entry Entry) {
	if a.Direction == entry.Direction {
		a.Balance += entry.Amount
		return
	}

	a.Balance -= entry.Amount
}
