package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	if got := Greeting("abc123"); got != "Hello from varo-sample-app abc123" {
		t.Fatalf("Greeting() = %q", got)
	}
}

func TestHealthCheck(t *testing.T) {
	rec := httptest.NewRecorder()
	NewHandler("abc123").ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/up", nil))
	if rec.Code != http.StatusOK || strings.TrimSpace(rec.Body.String()) != "ok" {
		t.Fatalf("/up = %d %q", rec.Code, rec.Body.String())
	}
}
