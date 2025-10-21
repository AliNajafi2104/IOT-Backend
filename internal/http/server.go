package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/IOT-Backend/internal/config"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, r *mux.Router, cfg *config.Config) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
