package models

import (
	"gorm.io/gorm"
	"time"
)

type UseragentService interface {
	Select(*Useragent) error
	InsertOrUpdate(*Useragent) error
}

type Useragent struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Useragent string         `json:"useragent" gorm:"unique"`
	Device    string         `json:"device"`
}
