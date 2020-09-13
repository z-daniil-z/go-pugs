package search

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/search/google"
	"gorm.io/gorm"
)

type API struct {
	db     *gorm.DB
	google *google.API
}

func NewAPI(db *gorm.DB) *API {
	return &API{db: db, google: google.NewAPI(db)}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/google", api.google.Router())
	return r
}
