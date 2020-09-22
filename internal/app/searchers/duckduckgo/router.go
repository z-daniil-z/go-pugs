package duckduckgo

import (
	"go-pugs/internal/db/postgres"
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type API struct {
	fileService models.FileService
}

func NewAPI(db *gorm.DB) *API {
	return &API{fileService: postgres.NewFileService(db)}
}
