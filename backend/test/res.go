package test

import "encoding/json"

// Status checks the status code of the response
func (t *Test) Status(expectedEngineStatus int) *Test {
	if t.rr.Code != expectedEngineStatus {
		t.t.Errorf("handler returned wrong status code: got %v want %v", t.rr.Code, expectedEngineStatus)
	}

	return t
}

// JSON checks the JSON response
func (t *Test) JSON() map[string]interface{} {
	if t.rr.Header().Get("Content-Type") != "application/json" {
		t.t.Errorf("handler returned wrong content type: got %v want %v", t.rr.Header().Get("Content-Type"), "application/json")
	}

	var response map[string]interface{}
	if err := json.Unmarshal(t.rr.Body.Bytes(), &response); err != nil {
		t.t.Errorf("handler returned invalid JSON: %v", err)
	}

	return response
}
