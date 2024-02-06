package api

import (
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

	// Default handlers
	r.NotFound(s.notFoundHandler)
	r.MethodNotAllowed(s.methodNotAllowedHandler)

	// Serve static files from the image store if it should serve static files
	if s.imageStore.ShouldServeStatic() {
		r.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir(s.imageStore.GetStaticPath()))))
	}

	// Register routes for v1
	r.Route("/v1", s.getV1Router)

	// Register routes for App API
	r.Route("/app", s.getAppRouter)

	// Default index handler
	r.Get("/", s.indexHandler)

	return r
}

// Register routes for v1
func (s *Server) getV1Router(r chi.Router) {
	// Public API routes
	r.Group(s.getPublicApiRouter)

	// Private API routes (API key authentication)
	r.Group(s.getPrivateApiRouter)
}
