package searchers

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/searchers/duckduckgo"
	"go-pugs/internal/app/searchers/google"
	"go-pugs/internal/middleware"
	"go-pugs/internal/usecases"
	"gorm.io/gorm"
)

type API struct {
	google     usecases.Search
	duckDuckGo usecases.Search
}

func NewAPI(db *gorm.DB) *API {
	return &API{google: google.NewAPI(db), duckDuckGo: duckduckgo.NewAPI(db)}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/google", api.Mount(api.google))
	r.Mount("/duckduckgo", api.Mount(api.duckDuckGo))
	return r
}

func (api *API) Mount(search usecases.Search) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", middleware.Context(search.SearchRequest).ServeHTTP)
	return r
}
