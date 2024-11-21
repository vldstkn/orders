package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"orders/internal/di"
	"orders/internal/domain"
	"orders/pkg/jwt"
	"time"
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

func (service *AccountService) Register(email, password, name string) (int, error) {
	existsUser := service.AccountRepository.FindByEmail(email)
	if existsUser != nil {
		return -1, errors.New("the user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	user := &domain.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}
	id, err := service.AccountRepository.Create(user)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (service *AccountService) Login(email, password string) (int, error) {
	user := service.AccountRepository.FindByEmail(email)
	if user == nil {
		return -1, nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return -1, errors.New("invalid email or password")
	}
	return user.Id, nil
}

func (service *AccountService) IssueTokens(secret string, data jwt.JWTData,
	expirationTimeA, expirationTimeR time.Time) (string, string, error) {

	j := jwt.NewJWT(secret)
	accessToken, err := j.Create(data, expirationTimeA)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := j.Create(data, expirationTimeR)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
