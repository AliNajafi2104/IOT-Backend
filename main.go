package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/IOT-Backend/config"
	"github.com/IOT-Backend/db"
	"github.com/IOT-Backend/handler"
	"github.com/IOT-Backend/mqtt"
	"github.com/IOT-Backend/repository"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, r *mux.Router, cfg *config.Config) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("starting server")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

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
		fx.Invoke(NewHTTPServer),
	).Run()
}
