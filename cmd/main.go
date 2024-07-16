package main

import (
	"log/slog"
	"net/http"

	"github.com/vgmrs/goexpert-weather-api/internal/infra/web"
)

func main() {
	r := web.Router()
	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	slog.Info("Start server!")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
