package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// Validate HTML
	if html == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("empty HTML"))
		return
	}
	// TODO: Validate HTML
	// if html == "" {
	// 	writeError(w, http.StatusBadRequest, fmt.Errorf("invalid HTML"))
	// 	return
	// }

	// Example http.html file
	if html == "http.html" {
		exampleHtml, err := os.ReadFile("../examples/http.html")

		if err != nil {
			writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot read example file"))
			return
		}

		html = string(exampleHtml)
	}

	// Example inline.html file
	if html == "inline.html" {
		exampleHtml, err := os.ReadFile("../examples/inline.html")

		if err != nil {
			writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot read example file"))
			return
		}

		html = string(exampleHtml)
	}

	// Grab a browser instance from the pool
	browser := s.browserPool.GetBrowser()

	// Render HTML
	image, err := browser.RenderHTML(html)

	if (err != nil) || (image == nil) {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot render HTML"))
		return
	}

	// Store image
	imageFilename := generateRandomString(10) + ".jpg"
	err = s.imageStore.StoreFile(imageFilename, image)

	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot store image"))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"time":  getElapsedtime(r).String(),
		"image": s.imageStore.GetFileURL(imageFilename),
	})
}

// Render URL handler
func (s *Server) renderUrlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	// Validate URL
	if !isValidUrl(url) {
		writeError(w, http.StatusBadRequest, fmt.Errorf("invalid URL"))
		return
	}

	// Grab a browser instance from the pool
	browser := s.browserPool.GetBrowser()

	// Render URL
	image, err := browser.RenderURL(url)

	if (err != nil) || (image == nil) {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot render URL"))
		return
	}

	// Store image
	imageFilename := generateRandomString(10) + ".jpg"
	err = s.imageStore.StoreFile(imageFilename, image)

	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("cannot store image"))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"time":  getElapsedtime(r).String(),
		"image": s.imageStore.GetFileURL(imageFilename),
	})
}
