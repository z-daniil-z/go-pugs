package search

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/search/google"
)

type API struct {
	google *google.API
}

func NewAPI() *API {
	return &API{google: google.NewAPI()}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/google", api.google.Router())
	return r
}
