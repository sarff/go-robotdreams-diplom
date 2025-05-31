package mongo

import (
	"context"

	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	DB *mongo.Database
}

func NewMongoClient(ctx context.Context, cfg *config.Config) (*DBClient, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.Database.URI)
	connection, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &DBClient{
		DB: connection.Database(cfg.Database.Name),
	}, nil
}
