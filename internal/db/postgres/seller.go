package postgres

import (
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type SellerService struct {
	db *gorm.DB
}

func NewSellerService(db *gorm.DB) *SellerService {
	ret := &SellerService{db: db}
	return ret
}

func (s *SellerService) Select(promo *models.Seller) error {
	return nil
}

func (s *SellerService) InsertOrUpdate(promo *models.Seller) error {
	return nil
}
