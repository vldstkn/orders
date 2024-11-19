package handlers

import (
	"context"
	pb "orders/pkg/api/account"
)

type AccountHandler struct {
	pb.UnimplementedAccountServer
}

type AccountGrpcHandlerDeps struct {
}

func NewAccountGrpcHandler() *AccountHandler {
	return &AccountHandler{}
}

func (handler *AccountHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Id:          0,
		AccessToken: req.Email + req.Name + req.Password,
	}, nil
}

func (handler *AccountHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Id:          0,
		AccessToken: req.Email + req.Password,
	}, nil
}
