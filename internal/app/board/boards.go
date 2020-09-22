package board

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type API struct {
}

func NewAPI(db *gorm.DB) *API {
	return &API{}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	return r
}
