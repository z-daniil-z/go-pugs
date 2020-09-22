package postgres

import (
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type PromoService struct {
	db *gorm.DB
}

func NewPromoService(db *gorm.DB) *PromoService {
	ret := &PromoService{db: db}
	return ret
}

func (p *PromoService) Select(promo *models.Promo) error {
	return nil
}

func (p *PromoService) InsertOrUpdate(promo *models.Promo) error {
	return nil
}
