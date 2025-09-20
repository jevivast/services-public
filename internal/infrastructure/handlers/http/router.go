package http

import (
	"github.com/gorilla/mux"
)

// Handler struct holds the HTTP handlers
type Handler struct {
	Router *mux.Router
}

// NewRouter creates a new router with all routes
func NewRouter() *mux.Router {
	h := &Handler{
		Router: mux.NewRouter(),
	}

	h.Router.HandleFunc("/api/v1/health", h.HealthCheck).Methods("GET")
	
	return h.Router
}
