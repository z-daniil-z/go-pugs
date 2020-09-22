package search

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/search/google"
	"go-pugs/internal/middleware"
	"go-pugs/internal/usecases"
	"gorm.io/gorm"
)

type API struct {
	google *google.API
}

func NewAPI(db *gorm.DB) *API {
	return &API{google: google.NewAPI(db)}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/google", api.Mount(api.google))
	return r
}

func (api *API) Mount(search usecases.Search) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", middleware.Context(search.SearchRequest).ServeHTTP)
	return r
}
