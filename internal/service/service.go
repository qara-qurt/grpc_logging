package service

import "github.com/qara-qurt/grpc_logging/internal/repository"

type ILogging interface {
	Insert()
}

type Service struct {
	Logging ILogging
}

func New(repo *repository.Repository) *Service {
	logging := newLogging(repo.Logging)
	return &Service{
		Logging: logging,
	}
}
