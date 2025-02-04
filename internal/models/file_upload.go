package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FileUpload Модель завантаженного файла.
type FileUpload struct {
	Uuid       string `gorm:"primaryKey;type:uuid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	EntityType string         `gorm:"column:entity_type;type:string;size:255"`
	EntityId   string         `gorm:"column:entity_id;type:string"`
	Name       string         `gorm:"column:name;type:string;size:255"`
	Path       string         `gorm:"column:path;type:string;size:255"`
}

// DTO Повертає DTO завантаженного файла.
func (f *FileUpload) DTO() FileUploadDto {
	return FileUploadDto{
		Uuid:       f.Uuid,
		EntityType: f.EntityType,
		EntityId:   f.EntityId,
		Name:       f.Name,
		Path:       f.Path,
	}
}

// BeforeCreate Генерація uuid.
func (n *FileUpload) BeforeCreate(tx *gorm.DB) (err error) {
	n.Uuid = uuid.New().String()
	return
}
