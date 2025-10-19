package repository

import (
	"context"

	"github.com/IOT-Backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) GetSites() ([]types.Site, error) {
	ctx := context.Background()

	coll := r.client.Database("iot").Collection("sites")
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

func (r *Repository) GetSiteById(id string) (*types.Site, error) {
	ctx := context.Background()

	filter := bson.D{{"id", id}}
	coll := r.client.Database("iot").Collection("sites")
	site := &types.Site{}
	err := coll.FindOne(ctx, filter).Decode(site)
	if err != nil {
		return nil, err
	}
	return site, nil
}
