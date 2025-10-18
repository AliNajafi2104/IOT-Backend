package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(mux *mux.Router) {

	mux.HandleFunc("/sites", getSites)
	mux.HandleFunc("/sites/{id}", getSitesById)
	mux.HandleFunc("/coordinaters/{id}", getCoordinatersById)
	mux.HandleFunc("/nodes/{id}", getNodesById)

}

func getSites(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
func getSitesById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func getCoordinatersById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func getNodesById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
