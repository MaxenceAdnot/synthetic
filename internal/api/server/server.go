package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/maxenceadnot/synthetic/internal/api/store"
)

type APIServer struct {
	Store store.URLCheckStore
	http.Handler
}

const jsonContentType = "application/json"

func NewAPIServer(store store.URLCheckStore) *APIServer {
	apiServer := new(APIServer)
	apiServer.Store = store

	router := http.NewServeMux()
	router.Handle("/checks/url/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			strID := r.URL.Path[len("/checks/url/"):]

			id, err := strconv.ParseUint(strID, 10, 32)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				return
			}

			sample, err := apiServer.Store.GetURLCheckByID(uint32(id))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)

				return
			}

			w.Header().Set("Content-Type", jsonContentType)
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(&sample); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Sorry, only GET method are supported.")
		}
	}))

	apiServer.Handler = router

	return apiServer
}
