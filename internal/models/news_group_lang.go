package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewsGroupLand Модель перекладу групин новин.
type NewsGroupLang struct {
	Uuid      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Rid       string         `gorm:"column:rid;type:string"`
	Loc       string         `gorm:"column:loc;type:string;size:5"`
	Title     string         `gorm:"column:title;type:string;size:255"`
}

// BeforeCreate Генерація uuid.
func (n *NewsGroupLang) BeforeCreate(tx *gorm.DB) (err error) {
	n.Uuid = uuid.New().String()
	return
}
