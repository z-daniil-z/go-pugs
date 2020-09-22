package app

import (
	"github.com/go-chi/chi"
	"go-pugs/internal/app/manager"
	"go-pugs/internal/app/search"
	"go-pugs/internal/middleware"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type API interface {
	Router() *chi.Mux
}

type APP struct {
	search  API
	manager API
}

func NewAPP(db *gorm.DB) (*APP, error) {
	ret := &APP{
		search:  search.NewAPI(db),
		manager: manager.NewAPI(db),
	}
	if err := db.AutoMigrate(models.File{}, models.Useragent{}, models.Proxy{}); err != nil {
		return nil, err
	}
	return ret, nil
}

func (app *APP) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Middleware)
	r.Mount("/search", app.search.Router())
	r.Mount("/manager", app.manager.Router())
	return r
}
