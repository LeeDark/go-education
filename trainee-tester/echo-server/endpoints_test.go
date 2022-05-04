package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// this type is sufficient for testing HTTP responses
func TestPing(t *testing.T) {
	mux := setEndpoints()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

// allows you to create an HTTP server for performing end-to-end HTTP tests
// in which you can send HTTP requests to the server using an HTTP client
func TestPingClient(t *testing.T) {
	mux := setEndpoints()

	server := httptest.NewServer(mux)
	defer server.Close()

	usersUrl := fmt.Sprintf("%s/ping", server.URL)
	request, _ := http.NewRequest("GET", usersUrl, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}

func BenchmarkPing(b *testing.B) {
	mux := setEndpoints()

	r, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, r)
	}
}
