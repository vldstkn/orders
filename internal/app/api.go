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
	httpRouter *http.ServeMux
	httpServer *http.Server
	Config     *configs.Config
	DB         *db.DB
	Logger     *slog.Logger
}

func NewApiApp(deps *ApiAppDeps) *ApiApp {

	router := http.NewServeMux()

	stack := middleware.Chain(
		middleware.HttpLogger(deps.Logger),
	)

	server := &http.Server{
		Addr:    deps.Config.ApiAddress,
		Handler: stack(router),
	}

	return &ApiApp{
		DB:         deps.DB,
		httpRouter: router,
		httpServer: server,
		Config:     deps.Config,
		Logger:     deps.Logger,
	}
}

func (app *ApiApp) Run() {
	apiService := services.NewApiService(app.Config.JWTSecret)

	// Grpc Clients
	accountConn, err := newClientConn(app.Config.AccountAddress)
	if err != nil {
		panic(err)
	}
	accountClient := pb.NewAccountClient(accountConn)

	// http handlers
	handlers.NewAccountHttpHandler(app.httpRouter, accountClient, apiService)
	handlers.NewOrdersHttpHandler(app.httpRouter)
	handlers.NewOfferingsHttpHandler(app.httpRouter)

	app.Logger.Info("Microservice API starts",
		slog.String("MODE", app.Config.Mode),
		slog.String("ADDRESS", app.Config.ApiAddress),
		slog.String("DNS", app.Config.Dsn),
	)

	app.httpServer.ListenAndServe()

}

func newClientConn(address string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
