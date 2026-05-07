package domain

type Entry struct {
	ID        string
	AccountID string
	Direction Direction
	Amount    int64
}
