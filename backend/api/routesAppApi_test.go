package api

import (
	"net/http"
	"testing"

	"github.com/paupenin/renderhook/backend/api/middleware"
	"github.com/paupenin/renderhook/backend/test"
)

func TestGetCurrentUserHandler(t *testing.T) {
	server := NewTestServer()
	test := test.NewTest(t)

	json := test.Handler(
		server.getCurrentUserHandler,
		middleware.StartTimeMiddleware,
	).Get("/app/user").Status(http.StatusOK).JSON()

	test.Assert(json["user"]).Equal("paupenin")
	test.Assert(json["time"]).NotEmpty()
}
