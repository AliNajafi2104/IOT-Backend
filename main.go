package main

import (
	"net/http"

	"github.com/IOT-Backend/db"
	"github.com/IOT-Backend/handler"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		handler.Module,
		fx.Provide(
			db.InitMongo,
			mux.NewRouter,
		),
		fx.Invoke(func(r *mux.Router) {
			http.ListenAndServe(":8000", r)
		}),
	).Run()
}
