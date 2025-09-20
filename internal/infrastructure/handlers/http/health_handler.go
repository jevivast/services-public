package http

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// HealthCheck handles the health check endpoint
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := healthResponse{
		Status:  "available",
		Message: "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(response)
}
