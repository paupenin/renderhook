package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// Req creates a new request
func (t *Test) Req(method string, url string, body io.Reader) *Test {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.t.Fatal(err)
	}

	t.rr = httptest.NewRecorder()

	t.handler.ServeHTTP(t.rr, req)

	return t
}

// Get creates a new GET request
func (t *Test) Get(url string) *Test {
	return t.Req("GET", url, nil)
}

// Post creates a new POST request
func (t *Test) Post(url string, body io.Reader) *Test {
	return t.Req("POST", url, body)
}

// Put creates a new PUT request
func (t *Test) Put(url string, body io.Reader) *Test {
	return t.Req("PUT", url, body)
}

// Delete creates a new DELETE request
func (t *Test) Delete(url string) *Test {
	return t.Req("DELETE", url, nil)
}
