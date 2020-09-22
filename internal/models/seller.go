package models

import "time"

type SellerService interface {
	Select(*Seller) error
	InsertOrUpdate(*Seller) error
}

type Seller struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Promos    []Promo   `json:"promos"`
}
