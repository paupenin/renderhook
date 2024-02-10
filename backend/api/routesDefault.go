package api

import (
	"fmt"
	"net/http"
)

// Default index handler
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"service": "renderhook",
		"time":    getElapsedtime(r).String(),
		"versions": map[string]string{
			"v1": s.config.GetURL() + "/v1",
		},
	})
}

// Default not found handler
func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// No logging here, no one cares about 404s
	writeError(w, http.StatusNotFound, fmt.Errorf("not found"))
}

// Default method not allowed handler
func (s *Server) methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	// No logging here, no one cares about 405s
	writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
}
