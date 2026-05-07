package repository

import "github.com/souz-dev/edger-system-go/internal/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	FindByID(id string) (*domain.Transaction, error)
}
