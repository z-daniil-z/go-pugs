package app

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/search"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type Router interface {
	Router() *chi.Mux
}

type APP struct {
	db     *gorm.DB
	search *search.API
}

func NewAPP(db *gorm.DB) (*APP, error) {
	ret := &APP{
		db:     db,
		search: search.NewAPI(),
	}
	if err := db.AutoMigrate(models.Files{}); err != nil {
		return nil, err
	}
	return ret, nil
}

func (app *APP) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/search", app.search.Router())
	return r
}
