package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module("repository",
	fx.Provide(NewRepository),
)

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

type Repository struct {
	client *mongo.Client
}
