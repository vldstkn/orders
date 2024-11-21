package services

import (
	"net/http"
	"time"
)

type ApiService struct {
	JWTSecret string
}

func NewApiService(JWTSecret string) *ApiService {
	return &ApiService{
		JWTSecret: JWTSecret,
	}
}

func (service *ApiService) AddCookie(w http.ResponseWriter, name, value string, expirationTime time.Time) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Secure:   false,
		Expires:  expirationTime,
		Path:     "/auth",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
