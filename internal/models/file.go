package models

import (
	"gorm.io/gorm"
	"time"
)

type FileService interface {
	Select(*File) error
	InsertOrUpdate(*File) error
}

type File struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Type      string         `json:"type"`
	Url       string         `json:"url" gorm:"unique"`
}
