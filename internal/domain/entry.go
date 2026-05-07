package domain

import "time"

type Entry struct {
	ID        string
	AccountID string
	Direction Direction
	Amount    int64
	CreatedAt time.Time
}
