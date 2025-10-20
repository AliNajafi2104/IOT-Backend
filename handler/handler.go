package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IOT-Backend/config"
	"github.com/IOT-Backend/repository"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(RegisterHandlers),
	fx.Invoke(NewHTTPServer),
)

type Params struct {
	fx.In

	Repo       repository.Repository
	MqttClient mqtt.Client
}

func NewHandler(p Params) *handler {
	return &handler{
		Repo:       p.Repo,
		mqttClient: p.MqttClient,
	}
}

type handler struct {
	Repo       repository.Repository
	mqttClient mqtt.Client
}

func NewHTTPServer(lc fx.Lifecycle, r *mux.Router, cfg *config.Config) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("starting server")
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
	router.HandleFunc("/ws", h.websocket)

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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *handler) websocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := ws.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(msgType, string(msg))

		err = ws.WriteMessage(msgType, msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
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
