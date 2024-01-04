package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Register routes for public API
func (s *Server) getPublicApiRouter(r chi.Router) {
	// r.Use(RateLimitMiddleware)

	r.Get("/status", s.healthStatusHandler)

	// TODO: Add more public routes
}

// Health Status handler
func (s *Server) healthStatusHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"time":    getElapsedtime(r).String(),
		"status":  "ok",
		"browser": "ready",
	})
}
