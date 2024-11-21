package handlers

import (
	"context"
	"orders/internal/di"
	pb "orders/pkg/api/account"
	"orders/pkg/jwt"
	"time"
)

type AccountHandler struct {
	pb.UnsafeAccountServer
	AccountService di.IAccountService
	JWTSecret      string
}

type AccountGrpcHandlerDeps struct {
	AccountService di.IAccountService
	JWTSecret      string
}

func NewAccountGrpcHandler(deps *AccountGrpcHandlerDeps) *AccountHandler {
	return &AccountHandler{
		AccountService: deps.AccountService,
		JWTSecret:      deps.JWTSecret,
	}
}

func (handler *AccountHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := handler.AccountService.Register(req.Email, req.Password, req.Name)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	accessToken, refreshToken, err := handler.AccountService.IssueTokens(handler.JWTSecret,
		jwt.JWTData{
			Id: id,
		},
		timeNow.Add(time.Hour*2),
		timeNow.AddDate(0, 0, 3))

	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Id:           int64(id),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (handler *AccountHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	id, err := handler.AccountService.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	accessToken, refreshToken, err := handler.AccountService.IssueTokens(handler.JWTSecret,
		jwt.JWTData{
			Id: id,
		},
		timeNow.Add(time.Hour*2),
		timeNow.AddDate(0, 0, 3))

	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Id:           int64(id),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
