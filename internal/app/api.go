package app

import (
	"log/slog"
	"net/http"
	"orders/internal/configs"
	handlers "orders/internal/handlers/http"
	"orders/internal/services"
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

	server := &http.Server{
		Addr:    deps.Config.Address,
		Handler: router,
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
	apiService := services.NewApiService()
	// handlers
	handlers.NewAccountHttpHandler(app.httpRouter, apiService)
	handlers.NewOrdersHttpHandler(app.httpRouter)

	app.Logger.Info("Server starts",
		slog.String("MODE", app.Config.Mode),
		slog.String("ADDRESS", app.Config.Address),
		slog.String("DNS", app.Config.Dsn),
	)

	app.httpServer.ListenAndServe()
}
