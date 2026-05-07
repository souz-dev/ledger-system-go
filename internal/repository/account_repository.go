package repository

import "github.com/souz-dev/edger-system-go/internal/domain"

type AccountRepository interface {
	Create(account *domain.Account) error
	FindByID(id string) (*domain.Account, error)
	Update(account *domain.Account) error
}
