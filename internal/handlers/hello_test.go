package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	expectedResponse := "чтобы найти самую длинную подстроку в строке, отправте ее POST запросом по url /api/substring"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response '%s', got '%s'", expectedResponse, recorder.Body.String())
	}
}
