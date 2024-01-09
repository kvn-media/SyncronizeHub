// utils/common_responses.go

package utils

import (
	"encoding/json"
	"net/http"
)

// CommonResponse represents a common JSON response structure.
type CommonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSONResponse writes a JSON response to the provided http.ResponseWriter.
func WriteJSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := CommonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
