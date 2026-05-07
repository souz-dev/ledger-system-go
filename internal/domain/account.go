package domain

type Account struct {
	ID        string
	Name      string
	Direction Direction
	Balance   int64
}

func (a *Account) Apply(entry Entry) {
	if a.Direction == entry.Direction {
		a.Balance += entry.Amount
		return
	}

	a.Balance -= entry.Amount
}
