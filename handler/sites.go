package handler

import (
	"encoding/json"
	"net/http"
)

func (h *handler) getSites(w http.ResponseWriter, r *http.Request) {
	sites, err := h.Repo.GetSites()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(sites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
func (h *handler) getSitesById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
