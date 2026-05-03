package main

import (
	"dift_backend_driver/driver-payments-service/internal/servicecore"
	"fmt"
	"log/slog"
	"net/http"

	"dift_backend_driver/driver-payments-service/config"
	"dift_backend_driver/driver-payments-service/internal/handler"
	"dift_backend_driver/driver-payments-service/internal/route"
	"dift_backend_driver/driver-payments-service/internal/service"
)

func main() {
	_ = servicecore.NewEngineUnifiedBundle(servicecore.LoadEngineUnifiedConfigFromEnv("driver-payments-service"))
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		panic(err)
	}
	logger := slog.Default()
	svc := service.NewPaymentsService()
	h := handler.NewPaymentsHandler(svc)
	mux := http.NewServeMux()
	route.Register(mux, h)
	addr := fmt.Sprintf(":%d", cfg.Server.HTTPPort)
	logger.Info("driver-payments-service started", "port", cfg.Server.HTTPPort)
	if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
