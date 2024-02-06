package test

import "testing"

// TestAssert is a test assertion
type TestAssert struct {
	t      *testing.T
	actual interface{}
}

// Assert checks the response
func (t *Test) Assert(actual interface{}) *TestAssert {
	return &TestAssert{
		t:      t.t,
		actual: actual,
	}
}

// Equal checks if the actual value is equal to the expected value
func (a *TestAssert) Equal(expected interface{}) {
	if a.actual != expected {
		a.t.Errorf("handler returned unexpected body: got %v want %v", a.actual, expected)
	}
}

// NotEmpty checks if the actual value is not empty
func (a *TestAssert) NotEmpty() {
	if a.actual == "" {
		a.t.Errorf("handler returned unexpected body: got %v want something", a.actual)
	}
}
