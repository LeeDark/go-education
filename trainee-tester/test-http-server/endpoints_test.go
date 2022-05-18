package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingSimple(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)

	ping(wr, req)

	//
	if wr.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "pong") {
		t.Errorf(
			`response body "%s" does not contain "NAME"`,
			wr.Body.String(),
		)
	}
}

func TestPingSub(t *testing.T) {
	t.Run("we run ping test as sub test", func(t *testing.T) {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/ping", nil)

		ping(wr, req)

		//
		if wr.Code != 200 {
			t.Errorf("HTTP Status expected: 200, got: %d", wr.Code)
		}

		if !strings.Contains(wr.Body.String(), "pong") {
			t.Errorf(
				`response body "%s" does not contain "NAME"`,
				wr.Body.String(),
			)
		}
	})
}

// this type is sufficient for testing HTTP responses
func TestPing(t *testing.T) {
	// 1
	mux := setEndpoints()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	// 2
	wr := httptest.NewRecorder()
	mux.ServeHTTP(wr, req)

	// 3
	if wr.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", wr.Code)
	}
}

// allows you to create an HTTP server for performing end-to-end HTTP tests
// in which you can send HTTP requests to the server using an HTTP client
func TestPingClient(t *testing.T) {
	// 1
	mux := setEndpoints()
	server := httptest.NewServer(mux)
	defer server.Close()

	usersUrl := fmt.Sprintf("%s/ping", server.URL)
	request, _ := http.NewRequest("GET", usersUrl, nil)

	// 2
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	// 3
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
