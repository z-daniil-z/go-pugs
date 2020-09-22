package models

import (
	"time"
)

type FileService interface {
	Select(*File) error
	InsertOrUpdate(*File) error
}

type File struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Type      string    `json:"type"`
	Url       string    `json:"url" gorm:"unique"`
}
