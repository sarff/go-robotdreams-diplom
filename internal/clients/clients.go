package clients

import (
	"context"

	appmongo "github.com/sarff/go-robotdreams-diplom/internal/clients/mongo"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
)

type Clients struct {
	Mongo *appmongo.DBClient
}

func NewClients(ctx context.Context, cfg *config.Config) (*Clients, error) {
	mongo, err := appmongo.NewMongoClient(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Clients{
		Mongo: mongo,
	}, nil
}
