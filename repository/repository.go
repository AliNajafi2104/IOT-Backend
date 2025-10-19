package repository

import (
	"context"

	"github.com/IOT-Backend/types"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *Repository) GetCoordinatorById(id string) (*types.Coordinator, error) {
	ctx := context.Background()

	filter := bson.D{{"id", id}}
	coll := r.client.Database("iot").Collection("coordinators")
	coordinator := &types.Coordinator{}
	err := coll.FindOne(ctx, filter).Decode(coordinator)
	if err != nil {
		return nil, err
	}
	return coordinator, nil
}

func (r *Repository) GetNodeById(id string) (*types.Node, error) {
	ctx := context.Background()

	filter := bson.D{{"id", id}}
	coll := r.client.Database("iot").Collection("nodes")
	node := &types.Node{}
	err := coll.FindOne(ctx, filter).Decode(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (r *Repository) GetOTAJobById(id string) (*types.OTAJob, error) {
	ctx := context.Background()

	filter := bson.D{{"id", id}}
	coll := r.client.Database("iot").Collection("OTAJob")
	OTAJob := &types.OTAJob{}
	err := coll.FindOne(ctx, filter).Decode(OTAJob)
	if err != nil {
		return nil, err
	}
	return OTAJob, nil
}
