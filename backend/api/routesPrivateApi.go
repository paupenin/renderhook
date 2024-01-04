package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) getPrivateApiRouter(r chi.Router) {
	// r.Use(APIKeyAuthMiddleware)

	// HTML routes
	r.Get("/html", s.renderHtmlHandler)
	r.Post("/html", s.renderHtmlHandler)

	// URL routes
	r.Get("/url", s.renderUrlHandler)
	r.Post("/url", s.renderUrlHandler)

	// TODO: Add more private routes
}

// Render HTML handler
func (s *Server) renderHtmlHandler(w http.ResponseWriter, r *http.Request) {
	html := r.FormValue("html")

	writeJSON(w, http.StatusOK, map[string]string{
		"time": getElapsedtime(r).String(),
		"html": html,
	})
}

// Render URL handler
func (s *Server) renderUrlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	writeJSON(w, http.StatusOK, map[string]string{
		"time": getElapsedtime(r).String(),
		"url":  url,
	})
}
