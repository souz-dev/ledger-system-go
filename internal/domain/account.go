package domain

import "time"

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Direction Direction `json:"direction"`
	Balance   int64     `json:"balance"` // Balance represents cents.
	CreatedAt time.Time `json:"created_at"`
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
