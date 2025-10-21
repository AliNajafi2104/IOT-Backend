package main

import (
	"github.com/IOT-Backend/config"
	"github.com/IOT-Backend/db"
	"github.com/IOT-Backend/http"
	"github.com/IOT-Backend/mqtt"
	"github.com/IOT-Backend/repository"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		mqtt.Module,
		http.Module,
		fx.Provide(
			zap.NewProduction,
			db.NewMongoDB,
			mux.NewRouter,
			config.LoadConfig,
			fx.Annotate(
				repository.NewMongoRepository,
				fx.As(new(repository.Repository)),
			),
		),
	).Run()
}
