package di

import "orders/internal/domain"

type IAccountService interface {
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
	FindByEmail(email string) *domain.User
}
type IAccountRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindByEmail(email string) *domain.User
}
