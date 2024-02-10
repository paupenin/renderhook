package api

import (
	"net/http"
	"testing"

	"github.com/paupenin/renderhook/backend/api/middleware"
	"github.com/paupenin/renderhook/backend/test"
)

func TestIndexHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	json := test.Handler(
		server.indexHandler,
		middleware.StartTimeMiddleware,
	).Get("/").Status(http.StatusOK).JSON()

	test.Assert(json["service"]).Equal("renderhook")
	test.Assert(json["time"]).NotEmpty()
	test.Assert(json["versions"]).NotEmpty()
	test.Assert(json["versions"].(map[string]interface{})["v1"]).Equal("/v1")
}

func TestNotFoundHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	test.Handler(
		server.notFoundHandler,
	).Get("/doesnotexist").Status(http.StatusNotFound)
}

func TestMethodNotAllowedHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	test.Handler(
		server.methodNotAllowedHandler,
	).Post("/", nil).Status(http.StatusMethodNotAllowed)
}
