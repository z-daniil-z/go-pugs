package postgres

import (
	"go-pugs/internal/models"
	"gorm.io/gorm"
)

type FileService struct {
	db *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	ret := &FileService{db: db}
	return ret
}

func (s *FileService) Insert(file *models.File) error {
	if err := s.db.Model(file).Where("url=?", file.Url).Update("url", file.Url).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err := s.db.Create(file).Error; err != nil {
		return err
	}
	return nil
}

func (s *FileService) Select(file *models.File) error {
	if file.Url != "" {
		if err := s.db.Model(file).Where("url=?", file.Url).Error; err != nil {
			return err
		}
	} else {
		if err := s.db.First(file, file.ID).Error; err != nil {
			return err
		}
	}
	return nil
}
