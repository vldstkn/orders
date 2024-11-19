package handlers

import (
	"context"
	"net/http"
	"orders/internal/di"
	"orders/internal/handlers/http/payload"
	pb "orders/pkg/api/account"
	"orders/pkg/req"
	"orders/pkg/res"
)

type AccountHttpHandler struct {
	ApiService    di.ApiService
	AccountClient pb.AccountClient
}

func NewAccountHttpHandler(router *http.ServeMux, accountClient pb.AccountClient, apiService di.ApiService) {
	handler := &AccountHttpHandler{
		ApiService:    apiService,
		AccountClient: accountClient,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /account/change_role", handler.ChangeRoleById())
	router.HandleFunc("PUT /account/{id}", handler.UpdateById())

	router.HandleFunc("GET /auth/login/access_token", handler.GetNewTokens())
	router.HandleFunc("GET /user/{id}", handler.GetPublicProfile())
}

func (handler *AccountHttpHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.AccountRegisterRequest](&w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		grpcRes, err := handler.AccountClient.Register(context.Background(), &pb.RegisterRequest{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		})

		res.Json(w, payload.AccountRegisterResponse{
			Email:       grpcRes.AccessToken,
			AccessToken: grpcRes.AccessToken,
		}, 201)
	}
}

func (handler *AccountHttpHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AccountHttpHandler) GetPublicProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AccountHttpHandler) GetNewTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AccountHttpHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AccountHttpHandler) ChangeRoleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
