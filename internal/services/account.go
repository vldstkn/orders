package services

import (
	"orders/internal/di"
	"orders/internal/domain"
)

type AccountService struct {
	AccountRepository di.IAccountRepository
}

type AccountServiceDeps struct {
	AccountRepository di.IAccountRepository
}

func NewAccountService(deps AccountServiceDeps) *AccountService {
	return &AccountService{
		AccountRepository: deps.AccountRepository,
	}
}

func (service *AccountService) Register(email, password string) (string, error) {
	return "", nil
}

func (service *AccountService) Login(email, password string) (string, error) {
	return "", nil
}

func (service *AccountService) FindByEmail(email string) *domain.User {
	return nil
}
