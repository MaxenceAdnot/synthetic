package store_test

import (
	"fmt"
	"testing"

	"github.com/maxenceadnot/synthetic/internal/api/store"
)

func TestURLCheckStore(t *testing.T) {
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

	t.Run("Create a new url check store", func(t *testing.T) {
		inmemStore := store.NewInMemoryURLCheckStore(urlChecks)

		if len(inmemStore.GetURLChecks()) != len(urlChecks) {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecks()), len(urlChecks))
		}

		if len(inmemStore.GetURLChecksByAccount(1)) != 2 {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecksByAccount(1)), 2)
		}

		if len(inmemStore.GetURLChecksByAccount(2)) != 1 {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecksByAccount(2)), 1)
		}

		if len(inmemStore.GetURLChecksByAccount(3)) != 0 {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecksByAccount(3)), 0)
		}

		if _, err := inmemStore.GetURLCheckByID(1); err != nil {
			t.Errorf("Got %v, want %v", err, nil)
		}

		if _, err := inmemStore.GetURLCheckByID(4); err == nil {
			t.Errorf("Got %v, want %v", err, fmt.Errorf("url check not found"))
		}

		if err := inmemStore.DeleteURLCheck(1); err != nil {
			t.Errorf("Got %v, want %v", err, nil)
		}

		if err := inmemStore.DeleteURLCheck(4); err == nil {
			t.Errorf("Got %v, want %v", err, fmt.Errorf("url check not found"))
		}

		if len(inmemStore.GetURLChecks()) != 2 {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecks()), 2)
		}

		newCheck, _ := store.NewURLCheck(1, "https://www.google.com", "GET", 200)

		if err := inmemStore.AddURLCheck(*newCheck); err != nil {
			t.Errorf("Got %v, want %v", err, nil)
		}

		if len(inmemStore.GetURLChecks()) != 3 {
			t.Errorf("Got %v, want %v", len(inmemStore.GetURLChecks()), 3)
		}
	})
}
