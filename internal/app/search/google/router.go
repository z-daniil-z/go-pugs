package google

import (
	"github.com/go-chi/chi"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.searchRequest)
	return r
}
