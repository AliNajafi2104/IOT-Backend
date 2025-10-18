package handler

import "net/http"

func (h *handler) getSites(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
func (h *handler) getSitesById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
