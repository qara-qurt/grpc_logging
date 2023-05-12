package repository

import (
	"errors"
	"github.com/qara-qurt/grpc_logging/configs"
	"github.com/qara-qurt/grpc_logging/internal/repository/mongo"
)

type ILogRepository interface {
	Insert()
}

type Repository struct {
	Logging ILogRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	mongoDB, err := mongo.NewDatabaseMongo(cfg)

	log := mongo.NewLogging(mongoDB.DB)
	if err != nil {
		return nil, errors.New("failed to connect mongoDB")
	}
	return &Repository{
		Logging: log,
	}, nil
}
