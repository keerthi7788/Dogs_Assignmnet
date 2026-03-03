package response

import (
	// Go internal packages
	"encoding/json"
	"net/http"
	// Local packages
)

// RespondJSON writes the data to the response writer as JSON
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if json.NewEncoder(w).Encode(data) != nil {
		http.Error(w, `{"message": "Internal Error Encoding Response"}`, http.StatusInternalServerError)
	}
}

// RespondMessage writes the message to the response writer
func RespondMessage(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"message": message})
}
