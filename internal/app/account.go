package app

import (
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"orders/internal/configs"
	handlers "orders/internal/handlers/grpc"
	"orders/internal/repositories"
	"orders/internal/services"
	pb "orders/pkg/api/account"
	"orders/pkg/db"
)

type AccountApp struct {
	Config *configs.Config
	DB     *db.DB
	Logger *slog.Logger
}

type AccountAppDeps struct {
	Config *configs.Config
	DB     *db.DB
	Logger *slog.Logger
}

func NewAccountApp(deps *AccountAppDeps) *AccountApp {
	return &AccountApp{
		DB:     deps.DB,
		Config: deps.Config,
		Logger: deps.Logger,
	}
}

func (app *AccountApp) Run() {
	accountRepository := repositories.NewAccountRepository(app.DB)
	accountService := services.NewAccountService(services.AccountServiceDeps{
		AccountRepository: accountRepository,
	})
	var opts []grpc.ServerOption
	accountHandler := handlers.NewAccountGrpcHandler(&handlers.AccountGrpcHandlerDeps{
		AccountService: accountService,
	})

	l, err := net.Listen("tcp", app.Config.AccountAddress)
	server := grpc.NewServer(opts...)
	pb.RegisterAccountServer(server, accountHandler)
	if err != nil {
		panic(err)
	}

	app.Logger.Info("Microservice Account starts",
		slog.String("MODE", app.Config.Mode),
		slog.String("ADDRESS", app.Config.AccountAddress),
		slog.String("DNS", app.Config.Dsn),
	)
	server.Serve(l)
}
