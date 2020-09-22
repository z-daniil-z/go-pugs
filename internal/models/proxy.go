package models

import (
	"gorm.io/gorm"
	"time"
)

type ProxyService interface {
	Select(*Proxy) error
	InsertOrUpdate(*Proxy) error
}

type Proxy struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Protocol  string         `json:"protocol"`
	Host      string         `json:"host" gorm:"unique"`
	Port      string         `json:"port"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}
