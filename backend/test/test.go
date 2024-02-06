package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test is a test
type Test struct {
	t       *testing.T
	handler http.HandlerFunc
	rr      *httptest.ResponseRecorder
}

// NewTest creates a new test
func NewTest(t *testing.T) *Test {
	return &Test{
		t: t,
	}
}
