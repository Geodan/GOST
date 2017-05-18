package rest

import (
	"net/http"

	"github.com/geodan/gost/sensorthings/models"
)

// handleDeleteRequest
func handleDeleteRequest(w http.ResponseWriter, e *models.Endpoint, r *http.Request, h *func() error) {
	handle := *h
	err := handle()
	if err != nil {
		sendError(w, []error{err})
		return
	}

	sendJSONResponse(w, http.StatusOK, nil, nil)
}
