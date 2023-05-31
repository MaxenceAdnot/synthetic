package store_test

import (
	"net/http"
	"testing"

	"github.com/maxenceadnot/synthetic/internal/api/store"
	"github.com/stretchr/testify/assert"
)

// Test that URLCheck can be created correctly using NewURLCheck.
func TestURLCheck(t *testing.T) {
	const urlTest = "https://www.google.com"

	getError := func(_ any, err error) error {
		return err
	}

	assert := assert.New(t)

	t.Run("Create a new url check with various url", func(t *testing.T) {
		url := urlTest
		method := http.MethodGet
		expectedStatus := 200

		assert.NoError(getError(store.NewURLCheck(0, url, method, expectedStatus)))
	})

	t.Run("Create a new url check with invalid url", func(t *testing.T) {
		url := "https://.invalid."
		method := http.MethodGet
		expectedStatus := 200

		assert.EqualError(getError(store.NewURLCheck(0, url, method, expectedStatus)), store.ErrInvalidURL.Error())
	})

	t.Run("Create a new url check with invalid method", func(t *testing.T) {
		url := urlTest
		method := "INVALID"
		expectedStatus := 200

		assert.EqualError(getError(store.NewURLCheck(0, url, method, expectedStatus)), store.ErrInvalidMethod.Error())
	})

	t.Run("Create a new url check with invalid expected status", func(t *testing.T) {
		url := urlTest
		method := "GET"
		expectedStatus := 999

		assert.EqualError(getError(store.NewURLCheck(0, url, method, expectedStatus)), store.ErrInvalidHTTPStatus.Error())
	})
}
