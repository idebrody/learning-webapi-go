package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type response struct {
	Result float64 `json:"result"`
}

func TestAdd(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/add?num1=3.14&num2=1.414", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { math(w, r, add) })
	handler.ServeHTTP(rr, req)

	// Check Status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the result
	var expected float64 = 4.554
	var resp response
	json.Unmarshal([]byte(rr.Body.String()), &resp)
	if expected != resp.Result {
		t.Errorf("Incorrect result: got %v want %v",
			resp.Result, expected)
	}
}
