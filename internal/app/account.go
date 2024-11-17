package app

import (
	"orders/internal/configs"
	"orders/internal/repositories"
	"orders/internal/services"
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
	accountRepository := repositories.NewAccountRepository(app.DB)
	accountService := services.NewAccountService(services.AccountServiceDeps{
		AccountRepository: accountRepository,
	})
	_ = accountService

}
