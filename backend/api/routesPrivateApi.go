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

	// Validate URL
	if !isValidUrl(url) {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid URL",
		})
		return
	}

	// Grab a browser instance from the pool
	browser := s.browserPool.GetBrowser()

	image, err := browser.RenderURL(url)

	if (err != nil) || (image == nil) {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Cannot render URL",
		})
		return
	}

	// Store image
	imageFilename := generateRandomString(10) + ".jpg"
	err = s.imageStore.StoreImage(imageFilename, image)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Cannot store image",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"time":  getElapsedtime(r).String(),
		"url":   url,
		"image": s.GetURL() + s.imageStore.GetImagePath(imageFilename),
	})
}
