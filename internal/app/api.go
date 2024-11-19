package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net/http"
	"orders/internal/configs"
	handlers "orders/internal/handlers/http"
	"orders/internal/handlers/middleware"
	"orders/internal/services"
	pb "orders/pkg/api/account"
	"orders/pkg/db"
)

type ApiAppDeps struct {
	Config *configs.Config
	DB     *db.DB
	Logger *slog.Logger
}

type ApiApp struct {
	httpRouter    *http.ServeMux
	httpServer    *http.Server
	Config        *configs.Config
	DB            *db.DB
	Logger        *slog.Logger
	AccountClient pb.AccountClient
}

func NewApiApp(deps *ApiAppDeps) *ApiApp {

	router := http.NewServeMux()

	stack := middleware.Chain(
		middleware.HttpLogger(deps.Logger),
	)

	server := &http.Server{
		Addr:    deps.Config.Address,
		Handler: stack(router),
	}
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:9876", opt)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	accountClient := pb.NewAccountClient(conn)

	return &ApiApp{
		DB:            deps.DB,
		httpRouter:    router,
		httpServer:    server,
		Config:        deps.Config,
		Logger:        deps.Logger,
		AccountClient: accountClient,
	}
}

func (app *ApiApp) Run() {
	apiService := services.NewApiService()

	// http handlers
	handlers.NewAccountHttpHandler(app.httpRouter, app.AccountClient, apiService)
	handlers.NewOrdersHttpHandler(app.httpRouter)
	handlers.NewOfferingsHttpHandler(app.httpRouter)

	app.Logger.Info("Server starts",
		slog.String("MODE", app.Config.Mode),
		slog.String("ADDRESS", app.Config.Address),
		slog.String("DNS", app.Config.Dsn),
	)

	app.httpServer.ListenAndServe()
}
