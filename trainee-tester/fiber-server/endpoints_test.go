package main

import (
	"net/http"
	"testing"
)

// this type is sufficient for testing HTTP responses
func TestPing(t *testing.T) {
	app := setEndpoints()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	resp, _ := app.Test(req, 1)
	if resp.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", resp.StatusCode)
	}
}

// allows you to create an HTTP server for performing end-to-end HTTP tests
// in which you can send HTTP requests to the server using an HTTP client
// func TestPingClient(t *testing.T) {
// 	app := setEndpoints()

// 	server := httptest.NewServer(app.Handler())
// 	defer server.Close()

// 	usersUrl := fmt.Sprintf("%s/ping", server.URL)
// 	request, _ := http.NewRequest("GET", usersUrl, nil)
// 	res, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if res.StatusCode != 200 {
// 		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
// 	}
// }

func BenchmarkPing(b *testing.B) {
	app := setEndpoints()

	r, _ := http.NewRequest("GET", "/ping", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		app.Test(r, 1)
	}
}
