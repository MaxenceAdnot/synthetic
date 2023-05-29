package api

import (
	"log"
	"net/http"

	"github.com/maxenceadnot/synthetic/internal/api/server"
	"github.com/maxenceadnot/synthetic/internal/api/store"
)

func Run() {
	log.Println("Starting API server")

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

	server := server.NewAPIServer(store.NewInMemoryURLCheckStore(urlChecks))

	log.Fatal(http.ListenAndServe(":8080", server))
}
