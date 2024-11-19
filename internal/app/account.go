package app

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"orders/internal/configs"
	handlers "orders/internal/handlers/grpc"
	pb "orders/pkg/api/account"
	"orders/pkg/db"
)

type AccountApp struct {
	Config *configs.Config
	DB     *db.DB
}

func NewAccountApp() *AccountApp {
	conf := configs.LoadConfig()
	database := db.NewDb(conf.Dsn)

	return &AccountApp{
		DB:     database,
		Config: conf,
	}
}

func (app *AccountApp) Run() {
	//accountRepository := repositories.NewAccountRepository(app.DB)
	//accountService := services.NewAccountService(services.AccountServiceDeps{
	//	AccountRepository: accountRepository,
	//})
	accountHandler := handlers.NewAccountGrpcHandler()
	server := grpc.NewServer()
	pb.RegisterAccountServer(server, accountHandler)

	l, err := net.Listen("tcp", ":9876")
	if err != nil {
		panic(err)
	}
	fmt.Println("Account Run :9876")
	server.Serve(l)
}
