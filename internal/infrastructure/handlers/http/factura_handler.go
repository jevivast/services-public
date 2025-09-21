package http

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

// GetFactura calls the external service with the provided idcontrato and proxies the response
func (h *Handler) GetFactura(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["idcontrato"]
	if id == "" {
		http.Error(w, "missing idcontrato", http.StatusBadRequest)
		return
	}

	base := h.BaseURL
	if base == "" {
		base = "https://empoduitama.com"
	}
	externalURL := base + "/page/insert_factura.php?idcontrato=" + url.QueryEscape(id)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, externalURL, nil)
	if err != nil {
		http.Error(w, "failed to build request", http.StatusInternalServerError)
		return
	}

	client := h.Client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "failed to reach external service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Proxy the status code and content-type
	if ct := resp.Header.Get("Content-Type"); ct != "" {
		w.Header().Set("Content-Type", ct)
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
