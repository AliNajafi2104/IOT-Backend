package repository

import (
	"context"

	"github.com/IOT-Backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repositoryImpl) GetSites() ([]types.Site, error) {
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

func (r *repositoryImpl) GetSiteById(id string) (*types.Site, error) {
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
