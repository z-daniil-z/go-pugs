package usecases

import (
	"go-pugs/internal/middleware"
	"net/http"
)

type Search interface {
	SearchRequest(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request)
}
