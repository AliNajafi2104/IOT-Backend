package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IOT-Backend/repository"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(RegisterHandlers),
)

func NewHandler(repo *repository.Repository, mqttClient mqtt.Client) *handler {
	return &handler{
		Repo:       repo,
		mqttClient: mqttClient,
	}
}

type handler struct {
	Repo       *repository.Repository
	mqttClient mqtt.Client
}

func RegisterHandlers(h *handler, router *mux.Router) {

	router.HandleFunc("/sites", h.getSites)
	router.HandleFunc("/sites/{id}", h.getSiteById)
	router.HandleFunc("/coordinaters/{id}", h.getCoordinatorById)
	router.HandleFunc("/nodes/{id}", h.getNodeById)
	router.HandleFunc("/color-profile", h.sendColorProfile)
	router.HandleFunc("/set-light", h.setLight)
	router.HandleFunc("/ota/start", h.StartOTAUpdate)
	router.HandleFunc("/ota/status", h.RetrieveOTAJobStatus)
	router.HandleFunc("/pairing/approve", h.ApproveNodePairing)

}

func (h *handler) getNodeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	node, err := h.Repo.GetNodeById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
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
