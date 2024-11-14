package app

import (
	"fmt"
	"net/http"
	"orders/internal/configs"
	handlers "orders/internal/handlers/http"
	"orders/internal/repositories"
	"orders/internal/services"
	"orders/pkg/db"
)

type AppAccount struct {
	httpRouter *http.ServeMux
	httpServer *http.Server
	Config     *configs.Config
	DB         *db.DB
}

func NewAccountApp() *AppAccount {
	conf := configs.LoadConfig()
	database := db.NewDb(conf.Dsn)

	router := http.NewServeMux()

	server := &http.Server{
		Addr:    conf.AccountAddress,
		Handler: router,
	}

	return &AppAccount{
		httpRouter: router,
		httpServer: server,
		DB:         database,
		Config:     conf,
	}
}

func (app *AppAccount) Run() {
	accountRepository := repositories.NewAccountRepository(app.DB)
	accountService := services.NewAccountService(services.AccountServiceDeps{
		AccountRepository: accountRepository,
	})
	handlers.NewAccountHandler(app.httpRouter, handlers.AccountHandlerDeps{
		AccountService: accountService,
	})

	fmt.Printf("server started on address: %s", app.Config.AccountAddress)
	err := app.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
