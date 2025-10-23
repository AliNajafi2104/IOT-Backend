package db

import (
	"context"
	"log"
	"time"

	"github.com/IOT-Backend/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type MongoDBResult struct {
	fx.Out

	MongoDB *mongo.Database
}

func NewMongoDB(cfg *config.Config) (MongoDBResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB.URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return MongoDBResult{}, err

	}
	db := client.Database("iot")
	return MongoDBResult{MongoDB: db}, nil
}
