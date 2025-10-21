package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/IOT-Backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ Repository = (*MongoRepository)(nil)

type Repository interface {
	GetCoordinatorById(id string) (*types.Coordinator, error)
	GetNodeById(id string) (*types.Node, error)
	GetOTAJobById(id string) (*types.OTAJob, error)
	GetSites() ([]types.Site, error)
	GetSiteById(id string) (*types.Site, error)
}

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db: db,
	}
}

func (r *MongoRepository) GetCoordinatorById(id string) (*types.Coordinator, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("coordinators")
	coordinator := &types.Coordinator{}
	err := coll.FindOne(ctx, filter).Decode(coordinator)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("coordinator with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return coordinator, nil
}

func (r *MongoRepository) GetNodeById(id string) (*types.Node, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("nodes")
	node := &types.Node{}
	err := coll.FindOne(ctx, filter).Decode(node)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("node with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return node, nil
}

func (r *MongoRepository) GetOTAJobById(id string) (*types.OTAJob, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("OTAJob")
	OTAJob := &types.OTAJob{}
	err := coll.FindOne(ctx, filter).Decode(OTAJob)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("OTA Job with id %s not found: %w", id, err)
		}
		return nil, err
	}
	return OTAJob, nil
}

func (r *MongoRepository) GetSites() ([]types.Site, error) {
	ctx := context.Background()

	coll := r.db.Collection("sites")
	cursor, err := coll.Find(ctx, struct{}{})
	if err != nil {
		return nil, err
	}

	results := []types.Site{}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetSiteById(id string) (*types.Site, error) {
	ctx := context.Background()

	filter := bson.D{{Key: "id", Value: id}}
	coll := r.db.Collection("sites")
	site := &types.Site{}
	err := coll.FindOne(ctx, filter).Decode(site)
	if err != nil {
		return nil, err
	}
	return site, nil
}
