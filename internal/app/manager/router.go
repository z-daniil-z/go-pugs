package manager

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/db/postgres"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type API struct {
	db          *gorm.DB
	fileService models.FileService
}

func NewAPI(db *gorm.DB) *API {
	return &API{
		db:          db,
		fileService: postgres.NewFileService(db),
	}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/file", api.getFile)
	return r
}
