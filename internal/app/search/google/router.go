package google

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/db/postgres"
	"go-pugs/internal/middleware"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type API struct {
	fileService models.FileService
}

func NewAPI(db *gorm.DB) *API {
	return &API{fileService: postgres.NewFileService(db)}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", middleware.Context(api.searchRequest).ServeHTTP)
	return r
}
