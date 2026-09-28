package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloTest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	rr := httptest.NewRecorder()
	HelloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)

		// 5. Cek Content-Type
		contentType := rr.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", contentType)
		}

		var res Response
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if res.Message != "Halo dari Backend Go (Golang)!" {
			t.Errorf("unexpected message: %s", res.Message)
		}

		if res.Status != "success" {
			t.Errorf("unexpected status: %s", res.Status)
		}

		if _, err := time.Parse(time.RFC3339, res.Timestamp); err != nil {
			t.Errorf("invalid timestamp format: %s", res.Timestamp)
		}
	}
}
