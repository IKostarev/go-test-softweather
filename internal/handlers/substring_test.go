package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubstringHandler(t *testing.T) {
	testCases := []struct {
		name       string
		method     string
		body       string
		statusCode int
		response   string
	}{
		{
			name:       "ValidRequest",
			method:     http.MethodPost,
			body:       "abcabcbb",
			statusCode: http.StatusOK,
			response:   "Исходная строка - abcabcbb, самая длинная подстрока без повторяющихся символов - abc\n",
		},
		{
			name:       "EmptyRequestBody",
			method:     http.MethodPost,
			body:       "",
			statusCode: http.StatusBadRequest,
			response:   "body is empty or ReadAll have error",
		},
		{
			name:       "InvalidRequestMethod",
			method:     http.MethodGet,
			body:       "abcabcbb",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/api/substring", bytes.NewBufferString(tc.body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(SubstringHandler)

			handler.ServeHTTP(recorder, req)

			if recorder.Code != tc.statusCode {
				t.Errorf("Expected status code %d, got %d", tc.statusCode, recorder.Code)
			}

			if tc.response != "" && recorder.Body.String() != tc.response {
				t.Errorf("Expected response '%s', got '%s'", tc.response, recorder.Body.String())
			}
		})
	}
}
