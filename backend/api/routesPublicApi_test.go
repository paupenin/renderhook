package api

import (
	"net/http"
	"testing"

	"github.com/paupenin/renderhook/backend/api/middleware"
	"github.com/paupenin/renderhook/backend/test"
)

func TestServiceInfoHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	json := test.Handler(
		server.serviceInfoHandler,
		middleware.StartTimeMiddleware,
	).Get("/v1").Status(http.StatusOK).JSON()

	test.Assert(json["service"]).Equal("renderhook")
	test.Assert(json["time"]).NotEmpty()
	test.Assert(json["version"]).Equal("1.0.0")
}

func TestHealthStatusHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	json := test.Handler(
		server.healthStatusHandler,
		middleware.StartTimeMiddleware,
	).Get("/v1/status").Status(http.StatusOK).JSON()

	test.Assert(json["time"]).NotEmpty()
	test.Assert(json["status"]).Equal("ok")
	test.Assert(json["engine"]).Equal("unavailable")

	// Initialize the browser pool
	server.browserPool.Init()

	// Do another request
	json = test.Handler(
		server.healthStatusHandler,
		middleware.StartTimeMiddleware,
	).Get("/status").Status(http.StatusOK).JSON()

	test.Assert(json["engine"]).Equal("ready")
}
