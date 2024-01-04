package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func isValidUrl(urlStr string) bool {
	// TODO: Force valid domain or valid IP
	// TODO: Disallow localhost and local IPs
	// TODO: Disallow private IPs
	// TODO: Disallow reserved IPs
	// TODO: Disallow reserved domains
	// TODO: Disallow reserved TLDs
	// TODO: Disallow reserved subdomains
	// TODO: Disallow reserved ports
	// TODO: Disallow reserved protocols

	// Parse the URL and ensure there's no error
	parsedUrl, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}

	// Check for the presence of a scheme (http/https) and a host
	return parsedUrl.Scheme != "" && parsedUrl.Host != ""
}

func generateRandomString(length int) string {
	// TODO: Generate a real random string
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
