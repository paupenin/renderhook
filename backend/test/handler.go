package test

import (
	"net/http"
)

// Handler sets the handler with optional middlewares
func (t *Test) Handler(handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) *Test {
	// Convert the handler to http.Handler type
	h := http.Handler(handler)

	// Apply middlewares in the order they are provided
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	t.handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})

	return t
}
