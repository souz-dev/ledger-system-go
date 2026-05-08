package domain

import (
	"errors"
	"testing"
	"time"
)

func TestTransactionValidate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		entries     []Entry
		expectedErr error
	}{
		{
			name:        "returns error when transaction has no entries",
			entries:     []Entry{},
			expectedErr: ErrTransactionWithoutEntries,
		},
		{
			name: "returns error when entry amount is invalid",
			entries: []Entry{
				{ID: "entry-1", AccountID: "acc-1", Direction: Debit, Amount: 0, CreatedAt: now},
				{ID: "entry-2", AccountID: "acc-2", Direction: Credit, Amount: 0, CreatedAt: now},
			},
			expectedErr: ErrInvalidEntryAmount,
		},
		{
			name: "returns error when entry direction is invalid",
			entries: []Entry{
				{ID: "entry-1", AccountID: "acc-1", Direction: Direction("invalid"), Amount: 100, CreatedAt: now},
				{ID: "entry-2", AccountID: "acc-2", Direction: Credit, Amount: 100, CreatedAt: now},
			},
			expectedErr: ErrInvalidDirection,
		},
		{
			name: "returns error when transaction is not balanced",
			entries: []Entry{
				{ID: "entry-1", AccountID: "acc-1", Direction: Debit, Amount: 100, CreatedAt: now},
				{ID: "entry-2", AccountID: "acc-2", Direction: Credit, Amount: 90, CreatedAt: now},
			},
			expectedErr: ErrTransactionNotBalanced,
		},
		{
			name: "valid transaction",
			entries: []Entry{
				{ID: "entry-1", AccountID: "acc-1", Direction: Debit, Amount: 100, CreatedAt: now},
				{ID: "entry-2", AccountID: "acc-2", Direction: Credit, Amount: 100, CreatedAt: now},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := &Transaction{
				ID:        "tx-1",
				Name:      "Test Transaction",
				Entries:   tt.entries,
				CreatedAt: now,
			}

			err := transaction.Validate()

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
