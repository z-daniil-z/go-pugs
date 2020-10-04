package boards

import (
	"github.com/go-chi/chi"
	auto_ru "go-pugs/internal/app/boards/auto.ru"
	"go-pugs/internal/middleware"
	"go-pugs/internal/usecases"
	"gorm.io/gorm"
)

type API struct {
	autoRu usecases.Board
}

func NewAPI(db *gorm.DB) *API {
	return &API{autoRu: auto_ru.NewAPI(db)}
}

func (api *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/autoru", api.Mount(api.autoRu))
	return r
}

func (api *API) Mount(board usecases.Board) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", middleware.Context(board.SearchRequest).ServeHTTP)
	return r
}
