package handlers

import (
	"context"
	"net/http"
	"orders/internal/configs"
	"orders/internal/di"
	"orders/internal/handlers/http/payload"
	"orders/internal/handlers/middleware"
	pb "orders/pkg/api/account"
	"orders/pkg/req"
	"orders/pkg/res"
	"time"
)

type AccountHttpHandler struct {
	ApiService    di.ApiService
	AccountClient pb.AccountClient
	Config        *configs.Config
}

type AccountHttpHandlerDeps struct {
	ApiService    di.ApiService
	AccountClient pb.AccountClient
	Config        *configs.Config
}

func NewAccountHttpHandler(router *http.ServeMux, deps *AccountHttpHandlerDeps) {
	handler := &AccountHttpHandler{
		ApiService:    deps.ApiService,
		AccountClient: deps.AccountClient,
		Config:        deps.Config,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /account/change_role", handler.ChangeRoleById())
	router.Handle("PUT  /account", middleware.IsAuthed(handler.UpdateById(), handler.Config))

	router.HandleFunc("GET /auth/login/access_token", handler.GetNewTokens())
	router.HandleFunc("GET /user/{id}", handler.GetPublicProfile())
}

func (handler *AccountHttpHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.AccountRegisterRequest](&w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		grpcRes, err := handler.AccountClient.Register(context.Background(), &pb.RegisterRequest{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		})

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		handler.ApiService.AddCookie(w, "refresh_token", grpcRes.RefreshToken, time.Now().AddDate(0, 0, 2))
		res.Json(w, payload.AccountRegisterResponse{
			Id:          int(grpcRes.Id),
			AccessToken: grpcRes.AccessToken,
		}, 201)
	}
}

func (handler *AccountHttpHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.AccountLoginRequest](&w, r)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		grpcRes, err := handler.AccountClient.Login(context.Background(), &pb.LoginRequest{
			Email:    body.Email,
			Password: body.Password,
		})

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		handler.ApiService.AddCookie(w, "refresh_token", grpcRes.RefreshToken, time.Now().AddDate(0, 0, 2))

		res.Json(w, payload.AccountLoginResponse{
			Id:          int(grpcRes.Id),
			AccessToken: grpcRes.AccessToken,
		}, http.StatusCreated)
	}
}

func (handler *AccountHttpHandler) GetPublicProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *AccountHttpHandler) GetNewTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		grpcRes, err := handler.AccountClient.GetNewTokens(context.Background(), &pb.GetNewTokensRequest{
			RefreshToken: cookie.Value,
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		handler.ApiService.AddCookie(w, "refresh_token", grpcRes.RefreshToken, time.Now().AddDate(0, 0, 2))
		res.Json(w, &payload.AccountGetNewTokensResponse{
			AccessToken: grpcRes.AccessToken,
		}, http.StatusOK)
	}
}

func (handler *AccountHttpHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(int)
		data, err := req.HandleBody[payload.AccountUpdateByIdRequest](&w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		protoReq := pb.UpdateUserRequest{
			Id: int64(id),
		}
		if data.NewPassword != nil {
			if data.Password == nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			protoReq.Password = data.Password
			protoReq.NewPassword = data.NewPassword
		}
		if data.Name != nil {
			protoReq.Name = data.Name
		}
		if data.Email != nil {
			protoReq.Email = data.Email
		}

		grpcRes, err := handler.AccountClient.UpdateById(context.Background(), &protoReq)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		res.Json(w, payload.AccountUpdateByIdResponse{
			IsSuccess: grpcRes.IsSuccess,
		}, http.StatusOK)
	}
}

func (handler *AccountHttpHandler) ChangeRoleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
