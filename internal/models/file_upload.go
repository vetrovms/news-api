package models

import (
	"gorm.io/gorm"
)

// FileUpload Модель завантаженного файла.
type FileUpload struct {
	gorm.Model
	EntityType string `gorm:"column:entity_type;type:string;size:255"`
	EntityId   int    `gorm:"column:entity_id;type:int"`
	Name       string `gorm:"column:name;type:string;size:255"`
	Path       string `gorm:"column:path;type:string;size:255"`
}

// DTO Повертає DTO завантаженного файла.
func (f *FileUpload) DTO() FileUploadDto {
	return FileUploadDto{
		ID:         int(f.ID),
		EntityType: f.EntityType,
		EntityId:   f.EntityId,
		Name:       f.Name,
		Path:       f.Path,
	}
}
