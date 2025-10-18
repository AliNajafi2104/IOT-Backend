package handler

import (
	"net/http"

	"github.com/IOT-Backend/repository"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(RegisterHandlers),
)

func NewHandler(router *mux.Router, repo *repository.Repository) *handler {
	return &handler{
		Router: router,
		Repo:   repo,
	}
}

type handler struct {
	Router *mux.Router
	Repo   *repository.Repository
}

func RegisterHandlers(h *handler) {

	h.Router.HandleFunc("/sites", h.getSites)
	h.Router.HandleFunc("/sites/{id}", h.getSitesById)
	h.Router.HandleFunc("/coordinaters/{id}", h.getCoordinatersById)
	h.Router.HandleFunc("/nodes/{id}", h.getNodesById)
	h.Router.HandleFunc("/color-profile", h.sendColorProfile)
	h.Router.HandleFunc("/set-light", h.setLight)
	h.Router.HandleFunc("/ota/start", h.StartOTAUpdate)
	h.Router.HandleFunc("/ota/status", h.RetrieveOTAJobStatus)
	h.Router.HandleFunc("/pairing/approve", h.ApproveNodePairing)

}

func (h *handler) getCoordinatersById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) getNodesById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) sendColorProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) setLight(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) StartOTAUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) RetrieveOTAJobStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *handler) ApproveNodePairing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
