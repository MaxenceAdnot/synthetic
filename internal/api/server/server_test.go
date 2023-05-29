package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/maxenceadnot/synthetic/internal/api/server"
	"github.com/maxenceadnot/synthetic/internal/api/store"
)

func TestGETURLCheck(t *testing.T) {
	urlChecks := []store.URLCheck{
		{
			ID:             1,
			AccountID:      1,
			URL:            "https://www.google.com",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			ID:             2,
			AccountID:      1,
			URL:            "https://www.example.com",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			ID:             3,
			AccountID:      2,
			URL:            "https://www.example.com",
			Method:         "GET",
			ExpectedStatus: 200,
		},
	}

	apiServer := server.NewAPIServer(store.NewInMemoryURLCheckStore(urlChecks))

	want, err := apiServer.Store.GetURLCheckByID(1)
	if err != nil {
		t.Fatalf("Unable to get url check by id %q, '%v'", 1, err)
	}

	t.Run("Get an URLCheck with a valid ID", func(t *testing.T) {
		id := 1

		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/checks/url/%d", id), nil)
		response := httptest.NewRecorder()

		apiServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		got := &store.URLCheck{}
		err := json.NewDecoder(response.Body).Decode(got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q, '%v'", response.Body, err)
		}

		if reflect.DeepEqual(got, want) != true {
			t.Errorf("Got %v, want %v", got, want)
		}
	})

	t.Run("Get an URLCheck a not existing ID", func(t *testing.T) {
		id := 4

		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/checks/url/%d", id), nil)
		response := httptest.NewRecorder()

		apiServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("Get an URLCheck a not existing ID", func(t *testing.T) {
		id := "test"

		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/checks/url/%s", id), nil)
		response := httptest.NewRecorder()

		apiServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("Get an URLCheck an invalid method", func(t *testing.T) {
		id := 1

		request, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("/checks/url/%d", id), nil)
		response := httptest.NewRecorder()

		apiServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusMethodNotAllowed)
	})
}

func assertStatus(tb testing.TB, got, want int) {
	tb.Helper()

	if got != want {
		tb.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
