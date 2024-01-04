package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) getAppRouter(r chi.Router) {
	// r.Use(AppAuthMiddleware)

	// User routes
	r.Route("/user", s.getUserRouter)
}

// Register routes for user
func (s *Server) getUserRouter(r chi.Router) {
	// Get current user
	r.Get("/", s.getCurrentUserHandler)
}

// Get current user handler
func (s *Server) getCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"time": getElapsedtime(r).String(),
		"user": "paupenin",
	})
}
