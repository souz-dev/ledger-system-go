package service

import (
	"testing"
	"time"

	"github.com/souz-dev/edger-system-go/internal/domain"
	"github.com/souz-dev/edger-system-go/internal/repository/memory"
)

func TestLedgerServiceCreateTransaction(t *testing.T) {
	now := time.Now()

	accountRepo := memory.NewAccountRepository()
	transactionRepo := memory.NewTransactionRepository()

	ledgerService := NewLedgerService(accountRepo, transactionRepo)

	debitAccount, err := domain.NewAccount("acc-1", "Cash", domain.Debit, 0, now)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	creditAccount, err := domain.NewAccount("acc-2", "Revenue", domain.Credit, 0, now)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := ledgerService.CreateAccount(debitAccount); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := ledgerService.CreateAccount(creditAccount); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	debitEntry, err := domain.NewEntry("entry-1", "acc-1", domain.Debit, 100, now)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	creditEntry, err := domain.NewEntry("entry-2", "acc-2", domain.Credit, 100, now)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	transaction, err := domain.NewTransaction(
		"tx-1",
		"Initial movement",
		[]domain.Entry{debitEntry, creditEntry},
		now,
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := ledgerService.CreateTransaction(transaction); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	updatedDebitAccount, err := ledgerService.GetAccountByID("acc-1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	updatedCreditAccount, err := ledgerService.GetAccountByID("acc-2")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedDebitAccount.Balance != 100 {
		t.Fatalf("expected debit account balance 100, got %d", updatedDebitAccount.Balance)
	}

	if updatedCreditAccount.Balance != 100 {
		t.Fatalf("expected credit account balance 100, got %d", updatedCreditAccount.Balance)
	}
}
