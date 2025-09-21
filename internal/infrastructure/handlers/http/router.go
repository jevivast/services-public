package http

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Handler struct holds the HTTP handlers
type Handler struct {
	Router  *mux.Router
	Client  *http.Client
	BaseURL string
}

// NewRouter creates a new router with all routes
func NewRouter() *mux.Router {
	h := &Handler{
		Router:  mux.NewRouter(),
		Client:  http.DefaultClient,
		BaseURL: "https://empoduitama.com",
	}

	h.Router.HandleFunc("/api/v1/health", h.HealthCheck).Methods("GET")
	h.Router.HandleFunc("/api/v1/factura/{idcontrato}", h.GetFactura).Methods("GET")

	// Serve OpenAPI YAML and Swagger UI
	// /api/v1/openapi.yaml -> serves file from ./api/v1/openapi.yaml
	h.Router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", http.FileServer(http.Dir("api"))))
	// /swagger/ -> serves Swagger UI from ./docs/swagger/
	h.Router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/swagger"))))
	
	return h.Router
}
