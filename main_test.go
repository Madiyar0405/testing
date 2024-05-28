package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestIndex tests the index handler.
func TestIndex(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := index()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedTitle := "Ex3_week3 | AI & GPT"
	if !contains(rr.Body.String(), expectedTitle) {
		t.Errorf("handler returned unexpected body: got %v want it to contain %v", rr.Body.String(), expectedTitle)
	}
}

// TestStartServer tests the server startup.
func TestStartServer(t *testing.T) {
	// Start the server in a separate goroutine.
	go func() {
		if err := startServer(); err != nil {
			t.Fatalf("could not start server: %v", err)
		}
	}()

	// Give the server a moment to start.
	time.Sleep(100 * time.Millisecond)

	// Make a request to the server.
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
}

// contains checks if a substring is present in a string.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
