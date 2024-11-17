package http

import (
	"fmt"
	"net/http"
	"orders/internal/di"
)

type AccountHttpHandler struct {
	ApiService di.ApiService
}

func NewAccountHttpHandler(router *http.ServeMux, apiService di.ApiService) {
	handler := &AccountHttpHandler{
		ApiService: apiService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("GET /auth/login/access_token", handler.GetNewTokens())
	router.HandleFunc("GET /user/{id}", handler.GetPublicProfile())
	router.HandleFunc("POST /account/change_role", handler.ChangeRole())
}

func (handler *AccountHttpHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func (handler *AccountHttpHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
	}
}

func (handler *AccountHttpHandler) GetPublicProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetPublicProfile")
	}
}

func (handler *AccountHttpHandler) GetNewTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetNewTokens")
	}
}

func (handler *AccountHttpHandler) ChangeRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetNewTokens")
	}
}
