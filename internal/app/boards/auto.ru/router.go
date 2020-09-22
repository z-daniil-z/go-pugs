package auto_ru

import (
	"go-pugs/internal/db/postgres"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type API struct {
	sellerService models.SellerService
	promoService  models.PromoService
}

func NewAPI(db *gorm.DB) *API {
	return &API{sellerService: postgres.NewSellerService(db), promoService: postgres.NewPromoService(db)}
}
