package repositories

import (
	"orders/internal/domain"
	"orders/pkg/db"
)

type AccountRepository struct {
	DB *db.DB
}

func NewAccountRepository(db *db.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (repo *AccountRepository) Create(user *domain.User) (*domain.User, error) {
	return nil, nil
}
func (repo *AccountRepository) FindByEmail(email string) *domain.User {
	return nil
}
