package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module("repository",
	fx.Provide(NewRepository),
)

func NewRepository(client *mongo.Client) *repository {
	return &repository{
		client: client,
	}
}

type repository struct {
	client *mongo.Client
}
