package memory

import (
	"errors"
	"sync"

	"github.com/souz-dev/edger-system-go/internal/domain"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountRepository struct {
	mu       sync.RWMutex
	accounts map[string]*domain.Account
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		accounts: make(map[string]*domain.Account),
	}
}

func (r *AccountRepository) Create(account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.accounts[account.ID] = account
	return nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, ok := r.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (r *AccountRepository) Update(account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.accounts[account.ID]; !ok {
		return ErrAccountNotFound
	}

	r.accounts[account.ID] = account
	return nil
}
