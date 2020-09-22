package usecases

import (
	"go-pugs/internal/middleware"
	"go-pugs/internal/models"
	"net/http"
)

type Announcement struct {
	Seller models.Seller
	Promo  models.Promo
}

type Board interface {
	GetAnnouncementInfo(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request)
	SearchRequest(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request)
}
