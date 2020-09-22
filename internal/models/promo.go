package models

import (
	"time"
)

type PromoService interface {
	Select(*Promo) error
	InsertOrUpdate(*Promo) error
}

type Promo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	SellerID  uint      `json:"sellerId"`
	Photo     string    `json:"photo"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Url       string    `json:"url"`
}
