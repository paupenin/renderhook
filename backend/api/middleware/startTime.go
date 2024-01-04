package middleware

import (
	"context"
	"net/http"
	"time"
)

type startTimeMiddlewareContextKey string

const StartTimeKey startTimeMiddlewareContextKey = "start_time"

// StartTimeMiddleware is a middleware that adds a start time to the request context
func StartTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), StartTimeKey, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
