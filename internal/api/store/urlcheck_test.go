package store_test

import (
	"net/http"
	"testing"

	"github.com/maxenceadnot/synthetic/internal/api/store"
)

// Test that URLCheck can be created correctly using NewURLCheck.
func TestURLCheck(t *testing.T) {
	const urlTest = "https://www.google.com"

	t.Run("Create a new url check with various url", func(t *testing.T) {
		url := urlTest
		method := http.MethodGet
		expectedStatus := 200

		_, err := store.NewURLCheck(0, url, method, expectedStatus)
		if err != nil {
			t.Errorf("Got %v, want %v", err, nil)
		}
	})

	t.Run("Create a new url check with invalid url", func(t *testing.T) {
		url := "https://.invalid."
		method := http.MethodGet
		expectedStatus := 200

		_, err := store.NewURLCheck(0, url, method, expectedStatus)

		if err == nil {
			t.Errorf("Got %v, want %v", err, "invalid url")
		}
	})

	t.Run("Create a new url check with invalid method", func(t *testing.T) {
		url := urlTest
		method := "INVALID"
		expectedStatus := 200

		_, err := store.NewURLCheck(0, url, method, expectedStatus)

		if err == nil {
			t.Errorf("Got %v, want %v", err, "invalid method")
		}
	})

	t.Run("Create a new url check with invalid expected status", func(t *testing.T) {
		url := urlTest
		method := "GET"
		expectedStatus := 999

		_, err := store.NewURLCheck(0, url, method, expectedStatus)

		if err == nil {
			t.Errorf("Got %v, want %v", err, "invalid expected status")
		}
	})
}
