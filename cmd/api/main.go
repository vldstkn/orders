package main

import (
	"log/slog"
	"net/http"
	"orders/internal/app/api"
	"orders/internal/configs"
	"orders/pkg/logger"
	"os"
)

func main() {
	config := configs.LoadConfig()
	logger := logger.NewLogger(os.Stdout)

	app := app.App(config, logger)

	logger.Info("Server starts",
		slog.String("MODE", config.Mode),
		slog.String("ADDRESS", config.Address),
		slog.String("DNS", config.Dsn),
	)
	server := http.Server{
		Addr:    config.Address,
		Handler: app,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
