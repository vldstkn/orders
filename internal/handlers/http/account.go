package handlers

import (
	"fmt"
	"net/http"
	"orders/internal/di"
	"orders/internal/handlers/http/payload"
	"orders/pkg/req"
	"orders/pkg/res"
)

type AccountHandler struct {
	AccountService di.IAccountService
}

type AccountHandlerDeps struct {
	AccountService di.IAccountService
}

func NewAccountHandler(router *http.ServeMux, deps AccountHandlerDeps) {
	handler := &AccountHandler{
		AccountService: deps.AccountService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("GET /login/access_token", handler.GetPublicProfile())
	router.HandleFunc("GET /user/{id}", handler.GetPublicProfile())
}

func (handler *AccountHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.AccountRegisterRequest](&w, r)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		email, err := handler.AccountService.Register(body.Email, body.Password)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		res.Json(w, payload.AccountRegisterResponse{
			Email: email,
		}, http.StatusCreated)
	}
}

func (handler *AccountHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login!")
	}
}

func (handler *AccountHandler) GetPublicProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetPublicProfile!")
	}
}

func (handler *AccountHandler) GetNewTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetNewTokens!")
	}
}
