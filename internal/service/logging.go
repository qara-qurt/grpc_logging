package service

import "github.com/qara-qurt/grpc_logging/internal/repository"

type Logging struct {
	repo repository.ILogRepository
}

func newLogging(repo repository.ILogRepository) *Logging {
	return &Logging{
		repo: repo,
	}
}

func (l Logging) Insert() {
	//TODO implement me
	panic("implement me")
}
