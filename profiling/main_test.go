package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{5, 10},
		{10, 45},
	}

	for _, test := range tests {
		result := compute(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d, but got %d", test.n, test.expected, result)
		}
	}
}

func TestHTTPServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug/pprof/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HTTP handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}
}

func TestMain(m *testing.M) {
	// Start the main function (server) in a separate goroutine
	go main()

	// Use a WaitGroup to wait for the server to start
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Allow the server to start
	}()

	// Wait for the server to start
	wg.Wait()

	// Run the tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
