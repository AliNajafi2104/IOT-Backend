package repository

import (
	"context"

	"github.com/IOT-Backend/types"
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
