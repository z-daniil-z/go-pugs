package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type API struct {
	client *http.Client
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Router() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(map[string]bool{"ok": true}); err != nil {

		}
	})
	return router
}
