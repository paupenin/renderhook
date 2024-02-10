package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Register routes for public API
func (s *Server) getPublicApiRouter(r chi.Router) {
	// r.Use(RateLimitMiddleware)

	// Service info
	r.Get("/", s.serviceInfoHandler)

	// Health status
	r.Get("/status", s.healthStatusHandler)
}

// Service info handler
func (s *Server) serviceInfoHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"service": "renderhook",
		"time":    getElapsedtime(r).String(),
		"version": "1.0.0",
	})
}

// Health Status handler
func (s *Server) healthStatusHandler(w http.ResponseWriter, r *http.Request) {
	engineStatus := "unavailable"

	if s.browserPool.IsReady() {
		engineStatus = "ready"
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"time":   getElapsedtime(r).String(),
		"status": "ok",
		"engine": engineStatus,
	})
}
