package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Greeting is the text of the home page.
func Greeting(version string) string {
	return fmt.Sprintf("Hi from varo-sample-app %s", version)
}

// NewHandler serves the app: / for people, /up for the health check,
// /version for the load test and /slow?ms= for long requests.
func NewHandler(version string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, Greeting(version))
	})
	mux.HandleFunc("GET /up", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, version)
	})
	mux.HandleFunc("GET /slow", func(w http.ResponseWriter, r *http.Request) {
		ms, _ := strconv.Atoi(r.URL.Query().Get("ms"))
		select {
		case <-time.After(time.Duration(min(ms, 30_000)) * time.Millisecond):
			_, _ = fmt.Fprintln(w, version)
		case <-r.Context().Done():
		}
	})
	return mux
}
