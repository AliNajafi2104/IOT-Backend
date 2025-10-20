package repository

import (
	"context"
	"fmt"

	"github.com/IOT-Backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module("repository",
	fx.Provide(NewRepository),
)

type Repository interface {
	GetCoordinatorById(id string) (*types.Coordinator, error)
	GetNodeById(id string) (*types.Node, error)
	GetOTAJobById(id string) (*types.OTAJob, error)
	GetSites() ([]types.Site, error)
	GetSiteById(id string) (*types.Site, error)
}

func NewRepository(db *mongo.Database) Repository {
	return &repositoryImpl{
		db: db,
	}
}

type repositoryImpl struct {
	db *mongo.Database
}

func (r *repositoryImpl) GetCoordinatorById(id string) (*types.Coordinator, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("coordinators")
	coordinator := &types.Coordinator{}
	err := coll.FindOne(ctx, filter).Decode(coordinator)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("coordinator with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return coordinator, nil
}

func (r *repositoryImpl) GetNodeById(id string) (*types.Node, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("nodes")
	node := &types.Node{}
	err := coll.FindOne(ctx, filter).Decode(node)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("node with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return node, nil
}

func (r *repositoryImpl) GetOTAJobById(id string) (*types.OTAJob, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("OTAJob")
	OTAJob := &types.OTAJob{}
	err := coll.FindOne(ctx, filter).Decode(OTAJob)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("OTA Job with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return OTAJob, nil
}
