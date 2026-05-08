package domain

import (
	"testing"
	"time"
)

func TestAccountApply(t *testing.T) {
	tests := []struct {
		name             string
		accountDirection Direction
		initialBalance   int64
		entryDirection   Direction
		entryAmount      int64
		expectedBalance  int64
	}{
		{
			name:             "increase balance when account and entry have same debit direction",
			accountDirection: Debit,
			initialBalance:   0,
			entryDirection:   Debit,
			entryAmount:      100,
			expectedBalance:  100,
		},
		{
			name:             "increase balance when account and entry have same credit direction",
			accountDirection: Credit,
			initialBalance:   0,
			entryDirection:   Credit,
			entryAmount:      100,
			expectedBalance:  100,
		},
		{
			name:             "decrease balance when debit account receives credit entry",
			accountDirection: Debit,
			initialBalance:   100,
			entryDirection:   Credit,
			entryAmount:      100,
			expectedBalance:  0,
		},
		{
			name:             "decrease balance when credit account receives debit entry",
			accountDirection: Credit,
			initialBalance:   100,
			entryDirection:   Debit,
			entryAmount:      100,
			expectedBalance:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := NewAccount("acc-1", "Test Account", tt.accountDirection, tt.initialBalance, time.Now())
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			entry, err := NewEntry("entry-1", account.ID, tt.entryDirection, tt.entryAmount, time.Now())
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			account.Apply(entry)

			if account.Balance != tt.expectedBalance {
				t.Fatalf("expected balance %d, got %d", tt.expectedBalance, account.Balance)
			}
		})
	}
}
