package mongo

import (
	"context"
	"github.com/qara-qurt/grpc_logging/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DatabaseMongo struct {
	DB *mongo.Database
}

func NewDatabaseMongo(cfg *configs.Config) (*DatabaseMongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	db := dbClient.Database(cfg.DB.Database)

	return &DatabaseMongo{
		DB: db,
	}, err
}
