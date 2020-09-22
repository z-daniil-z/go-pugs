package models

import (
	"time"
)

type Seller struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Promos    []Promo
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
