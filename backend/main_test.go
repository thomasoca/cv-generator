package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMainApiGet(t *testing.T) {

	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getExample)

	handler.ServeHTTP(rr, req)
	result := rr.Result()

	// Check the status code is what we expect.
	if status := result.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v with error %v",
			status, http.StatusOK, rr.Body.String())
	}

	// Check the header
	expectedHeader := "application/json"

	if result.Header.Get("Content-type") != expectedHeader {
		t.Errorf("handler returned unexpected header: got %v want %v",
			result.Header.Get("Content-type"), expectedHeader)
	}

}

func TestMainApi(t *testing.T) {
	body := JsonInput("./examples/user.json")
	req, err := http.NewRequest("POST", "/user", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveFile)

	handler.ServeHTTP(rr, req)
	result := rr.Result()

	// Check the status code is what we expect.
	if status := result.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v with error %v",
			status, http.StatusOK, rr.Body.String())
	}

	// Check the header
	expectedHeader := "application/pdf"

	if result.Header.Get("Content-type") != expectedHeader {
		t.Errorf("handler returned unexpected header: got %v want %v",
			result.Header.Get("Content-type"), expectedHeader)
	}

	// Check the leftover file
	_, e := os.Stat("./test/test.tex")
	if e == nil {
		t.Errorf("handler failed to processing leftover file")
	}

	req.Body.Close()
}

// Test API when faulty input used
func TestErrorRequestApi(t *testing.T) {
	body, _ := json.Marshal("")
	req, err := http.NewRequest("POST", "/user", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveFile)

	handler.ServeHTTP(rr, req)
	result := rr.Result()

	// Check the status code is what we expect.
	if status := result.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v with error %v",
			status, http.StatusBadRequest, rr.Body.String())
	}

	// Check the header
	expectedHeader := "application/json"

	if result.Header.Get("Content-type") != expectedHeader {
		t.Errorf("handler returned unexpected header: got %v want %v",
			result.Header.Get("Content-type"), expectedHeader)
	}

	// Check response body
	expectedMessage := `{"message": "Bad request"}`

	if rr.Body.String() != expectedMessage {
		t.Errorf("handler returned unexpected message: got '%v' want '%v'", rr.Body.String(), expectedMessage)
	}

	// Check the leftover file
	_, e := os.Stat("./test/test.tex")
	if e == nil {
		t.Errorf("handler failed to processing leftover file")
	}

	req.Body.Close()
}

// Test API when creating tex file failed
func TestErrorServerApi(t *testing.T) {
	body := JsonInput("./examples/user_error.json")
	req, err := http.NewRequest("POST", "/user", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveFile)

	handler.ServeHTTP(rr, req)
	result := rr.Result()

	// Check the status code is what we expect.
	if status := result.StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v with error %v",
			status, http.StatusInternalServerError, rr.Body.String())
	}

	// Check the header
	expectedHeader := "application/json"

	if result.Header.Get("Content-type") != expectedHeader {
		t.Errorf("handler returned unexpected header: got %v want %v",
			result.Header.Get("Content-type"), expectedHeader)
	}

	// Check response body
	expectedMessage := `{"message": "Failed creating file"}`

	if rr.Body.String() != expectedMessage {
		t.Errorf("handler returned unexpected message: got '%v' want '%v'", rr.Body.String(), expectedMessage)
	}

	// Check the leftover file
	_, e := os.Stat("./test/test.tex")
	if e == nil {
		t.Errorf("handler failed to processing leftover file")
	}

	req.Body.Close()
}
