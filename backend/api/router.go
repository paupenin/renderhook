package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paupenin/web2image/backend/api/middleware"
)

func (s *Server) initRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	// TODO: Add middlewares

	// Add StartTimeMiddleware
	r.Use(middleware.StartTimeMiddleware)

	// NotFound handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// No logging here, no one cares about 404s
		writeError(w, http.StatusNotFound, fmt.Errorf("not found"))
	})

	// MethodNotAllowed handler
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		// No logging here, no one cares about 405s
		writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	})

	// Register routes for v1
	r.Route("/v1", s.getV1Router)

	// Register routes for App API
	r.Route("/app", s.getAppRouter)

	return r
}

// Register routes for v1
func (s *Server) getV1Router(r chi.Router) {
	// Public API routes
	r.Group(s.getPublicApiRouter)

	// Private API routes (API key authentication)
	r.Group(s.getPrivateApiRouter)
}
