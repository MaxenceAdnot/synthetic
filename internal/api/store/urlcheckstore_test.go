package store_test

import (
	"testing"

	"github.com/maxenceadnot/synthetic/internal/api/store"
	"github.com/stretchr/testify/assert"
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

	getError := func(_ any, err error) error {
		return err
	}

	assert := assert.New(t)

	t.Run("Create a new url check store", func(t *testing.T) {
		inmemStore := store.NewInMemoryURLCheckStore(urlChecks)

		assert.Len(inmemStore.GetURLChecks(), len(urlChecks))
		assert.Len(inmemStore.GetURLChecksByAccount(1), 2)
		assert.Len(inmemStore.GetURLChecksByAccount(2), 1)
		assert.Len(inmemStore.GetURLChecksByAccount(3), 0)

		assert.NoError(getError(inmemStore.GetURLCheckByID(1)))

		assert.EqualError(getError(inmemStore.GetURLCheckByID(4)), store.ErrURLCheckNotFound.Error())

		assert.NoError(inmemStore.DeleteURLCheck(1))
		assert.EqualError(inmemStore.DeleteURLCheck(4), store.ErrURLCheckNotFound.Error())

		assert.Len(inmemStore.GetURLChecks(), len(urlChecks)-1)

		newCheck, _ := store.NewURLCheck(1, "https://www.google.com", "GET", 200)
		assert.NoError(inmemStore.AddURLCheck(*newCheck))
		assert.Len(inmemStore.GetURLChecks(), len(urlChecks))
	})
}
