package http

import (
	"net/http"

	"github.com/travisbale/go-template/sdk"
)

// HandleHealth returns the service health status
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	response := sdk.HealthResponse{
		Status: "OK",
	}

	respondJSON(w, http.StatusOK, response)
}
