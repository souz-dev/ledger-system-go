package service

import (
	"sync"

	"github.com/souz-dev/edger-system-go/internal/domain"
	"github.com/souz-dev/edger-system-go/internal/repository"
)

type LedgerService struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
	mu              sync.Mutex
}

func NewLedgerService(
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
) *LedgerService {
	return &LedgerService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *LedgerService) CreateTransaction(transaction *domain.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := transaction.Validate(); err != nil {
		return err
	}

	accounts := make(map[string]*domain.Account)

	for _, entry := range transaction.Entries {
		account, err := s.accountRepo.FindByID(entry.AccountID)
		if err != nil {
			return err
		}

		accounts[entry.AccountID] = account
	}

	for _, entry := range transaction.Entries {
		account := accounts[entry.AccountID]
		account.Apply(entry)

		if err := s.accountRepo.Update(account); err != nil {
			return err
		}
	}

	if err := s.transactionRepo.Create(transaction); err != nil {
		return err
	}

	return nil
}
