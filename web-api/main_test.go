package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{"Home Handler", home, http.StatusOK, "Welcome to the Shapes API"},
		{"Health Handler", health, http.StatusOK, "Server is running"},
		{"About Handler", about, http.StatusOK, "SHU-WEI"},
		{"Time Handler", currentTime, http.StatusOK, "Current server time"},
		{"Greeting Handler", greeting, http.StatusOK, "Hello from SHU-WEI’s Shapes API!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got %v, expected %v", rr.Code, tt.expectedStatus)
			}
			if !contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("got %v, expected to contain %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

// helper function
func contains(body, expected string) bool {
	return len(body) >= len(expected) && body[:len(expected)] == expected
}
