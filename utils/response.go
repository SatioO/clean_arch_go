package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, _ := json.Marshal(map[string]interface{}{
		"data":       payload,
		"statusCode": code,
	})

	w.WriteHeader(code)
	w.Write(response)
}
