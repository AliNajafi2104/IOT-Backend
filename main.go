package main

import (
	"github.com/IOT-Backend/config"
	"github.com/IOT-Backend/db"
	"github.com/IOT-Backend/handler"
	"github.com/IOT-Backend/mqtt"
	"github.com/IOT-Backend/repository"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		mqtt.Module,
		repository.Module,
		handler.Module,
		fx.Provide(
			zap.NewProduction,
			db.InitMongo,
			mux.NewRouter,
			config.LoadConfig,
		),
	).Run()
}
