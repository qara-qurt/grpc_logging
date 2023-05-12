package main

import (
	"errors"
	"github.com/qara-qurt/grpc_logging/configs" //nolint:typecheck
	"github.com/qara-qurt/grpc_logging/internal/repository"
	"github.com/qara-qurt/grpc_logging/internal/server"
	"github.com/qara-qurt/grpc_logging/internal/service"
	"log"
)

func main() {
	log.Fatal(run())
}

func run() error {

	cfg, err := configs.NewConfig()
	if err != nil {
		return errors.New("failed to get config")
	}

	repo, err := repository.New(cfg)
	if err != nil {
		return err
	}
	service := service.New(repo)
	loggingSrv := server.NewLoggingServer(service)
	srv := server.New(loggingSrv)

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		return err
	}
	return nil
}
