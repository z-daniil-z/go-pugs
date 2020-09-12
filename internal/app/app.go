package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"net/http"
)

type APP struct {
	db *gorm.DB
}

func NewAPP(db *gorm.DB) (*APP, error) {
	ret := &APP{db: db}
	return ret, nil
}

func (app *APP) Router() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(map[string]bool{"ok": true}); err != nil {

		}
	})
	return router
}
