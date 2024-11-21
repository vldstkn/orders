package di

import (
	"orders/internal/domain"
	"orders/pkg/jwt"
	"time"
)

type IAccountService interface {
	Register(email, password, name string) (int, error)
	Login(email, password string) (int, error)
	IssueTokens(secret string, data jwt.JWTData,
		expirationTimeA, expirationTimeR time.Time) (string, string, error)
}
type IAccountRepository interface {
	Create(user *domain.User) (int, error)
	FindByEmail(email string) *domain.User
}
