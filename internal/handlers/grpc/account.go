package handlers

import (
	"context"
	"errors"
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

func (handler *AccountHandler) GetNewTokens(ctx context.Context, req *pb.GetNewTokensRequest) (*pb.GetNewTokensResponse, error) {
	isValid, data := jwt.NewJWT(handler.JWTSecret).Parse(req.RefreshToken)
	if !isValid {
		return nil, errors.New("token is not valid")
	}
	timeNow := time.Now()
	accessToken, refreshToken, err := handler.AccountService.IssueTokens(handler.JWTSecret, *data,
		timeNow.Add(time.Hour*2),
		timeNow.AddDate(0, 0, 3))

	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &pb.GetNewTokensResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (handler *AccountHandler) UpdateById(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// TODO:
	return &pb.UpdateUserResponse{
		IsSuccess: false,
	}, nil
}
