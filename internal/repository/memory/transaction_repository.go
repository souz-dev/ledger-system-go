package memory

import (
	"errors"
	"sync"

	"github.com/souz-dev/edger-system-go/internal/domain"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type TransactionRepository struct {
	mu           sync.RWMutex
	transactions map[string]*domain.Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		transactions: make(map[string]*domain.Transaction),
	}
}

func (r *TransactionRepository) Create(transaction *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.transactions[transaction.ID] = transaction
	return nil
}

func (r *TransactionRepository) FindByID(id string) (*domain.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transaction, ok := r.transactions[id]
	if !ok {
		return nil, ErrTransactionNotFound
	}

	return transaction, nil
}
