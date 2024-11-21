package main

import (
	"orders/internal/app"
	"orders/internal/configs"
	"orders/pkg/db"
	"orders/pkg/logger"
	"os"
)

func main() {

	conf := configs.LoadConfig()
	logger := logger.NewLogger(os.Stdout)
	database := db.NewDb(conf.Dsn)

	accountApp := app.NewAccountApp(&app.AccountAppDeps{
		Config: conf,
		Logger: logger,
		DB:     database,
	})
	accountApp.Run()
}
