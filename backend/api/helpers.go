package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/paupenin/web2image/backend/api/middleware"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ErrorResponse struct {
	Status int    `json:"code"`
	Error  string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, ErrorResponse{
		Status: status,
		Error:  err.Error(),
	})
}

func getStarttime(r *http.Request) time.Time {
	return r.Context().Value(middleware.StartTimeKey).(time.Time)
}

func getElapsedtime(r *http.Request) time.Duration {
	return time.Since(getStarttime(r))
}
