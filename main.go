package main

import (
	"net/http"

	"github.com/IOT-Backend/db"
	"github.com/IOT-Backend/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			db.InitMongo,
			mux.NewRouter,
		),
		fx.Invoke(func(r *mux.Router) {
			handlers.RegisterHandlers(r)
			http.ListenAndServe(":8000", r)
		}),
	).Run()
}
