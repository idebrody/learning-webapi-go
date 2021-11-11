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

type errResponse struct {
	Error string `json:"error"`
}

func performRequest(h http.HandlerFunc, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	return rr
}

func assertEqual(expected float64, received float64, t *testing.T) {
	if expected != received {
		t.Errorf("Incorrect result: got %v want %v",
			received, expected)
	}
}

func assertError(expected string, received string, t *testing.T) {
	if expected != received {
		t.Errorf("Incorrect error: got '%v' want '%v'",
			received, expected)
	}
}

func assertStatus(expected int, received int, t *testing.T) {
	if expected != received {
		t.Errorf("Incorret status code: got %v want %v",
			received, expected)
	}
}

func assertArrayLength(array []int, expectedLength int, t *testing.T) {
	if len(array) != expectedLength {
		t.Errorf("Incorret array length: got %v want %v",
			len(array), expectedLength)
	}
}

func TestAdd(t *testing.T) {
	// Add handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { math(w, r, add) })

	t.Run("Adding Two Numbers Together", func(t *testing.T) {
		rr := performRequest(handler, "/add?num1=1.726&num2=1.414")
		wantStatus := http.StatusOK
		var want float64 = 3.14

		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp response
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertEqual(want, resp.Result, t)
	})

	t.Run("Check Parsing Error Reported Num1", func(t *testing.T) {
		rr := performRequest(handler, "/add?num1=3.14a&num2=1.414")
		wantStatus := http.StatusBadRequest
		want := "Invalid number: 3.14a"

		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp errResponse
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertError(want, resp.Error, t)
	})

	t.Run("Check Parsing Error Reported Num2", func(t *testing.T) {
		rr := performRequest(handler, "/add?num1=3.14&num2=--1.414")
		wantStatus := http.StatusBadRequest
		wantError := "Invalid number: --1.414"

		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp errResponse
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertError(wantError, resp.Error, t)
	})
}

func TestSubstract(t *testing.T) {
	// Add handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { math(w, r, substract) })

	t.Run("Substract Two Numbers", func(t *testing.T) {
		rr := performRequest(handler, "/substract?num1=3.14&num2=1.414")
		wantStatus := http.StatusOK
		var want float64 = 1.726

		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp response
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertEqual(want, resp.Result, t)
	})

}

func TestDivision(t *testing.T) {
	// Add handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { math(w, r, division) })

	t.Run("Divide Two Numbers", func(t *testing.T) {
		rr := performRequest(handler, "/division?num1=4.43996&num2=1.414")
		wantStatus := http.StatusOK
		var want float64 = 3.14
		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp response
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertEqual(want, resp.Result, t)
	})

	t.Run("Divede by Zero", func(t *testing.T) {
		rr := performRequest(handler, "/division?num1=4.43996&num2=0")
		wantStatus := http.StatusBadRequest
		wantError := "Division by Zero"

		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

		// Check the result
		var resp errResponse
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertError(wantError, resp.Error, t)
	})

}

func TestHealhChecks(t *testing.T) {
	// Add handler
	handler := http.HandlerFunc(liveness)

	t.Run("Liveness", func(t *testing.T) {
		rr := performRequest(handler, "/liveness")
		wantStatus := http.StatusOK
		// Check Status code
		assertStatus(wantStatus, rr.Code, t)

	})

}

func TestRandom(t *testing.T) {
	// Add handler
	handler := http.HandlerFunc(random)

	t.Run("Random no Count defined", func(t *testing.T) {
		rr := performRequest(handler, "/random")
		wantStatus := http.StatusOK
		wantLength := 10
		// Check Status code
		assertStatus(wantStatus, rr.Code, t)
		// Check result
		var resp []int
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertArrayLength(resp, wantLength, t)

	})

	t.Run("Random custom Count", func(t *testing.T) {
		rr := performRequest(handler, "/random?Count=25000")
		wantStatus := http.StatusOK
		wantLength := 25000
		// Check Status code
		assertStatus(wantStatus, rr.Code, t)
		// Check result
		var resp []int
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertArrayLength(resp, wantLength, t)

	})

	t.Run("Random Negative Count", func(t *testing.T) {
		rr := performRequest(handler, "/random?Count=-13")
		wantStatus := http.StatusBadRequest
		wantError := "Count must be positive: -13"
		// Check Status code
		assertStatus(wantStatus, rr.Code, t)
		// Check result
		var resp errResponse
		json.Unmarshal([]byte(rr.Body.String()), &resp)
		assertError(wantError, resp.Error, t)

	})

}
