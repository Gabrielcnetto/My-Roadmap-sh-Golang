package basics

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicExample(t *testing.T) {
	// create a new test server with a simple handler
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"message": "hello world"}`)
	}))
	defer server.Close()

	// make an HTTP request to the test server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// check the response
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %v", resp.Status)
	}
}
