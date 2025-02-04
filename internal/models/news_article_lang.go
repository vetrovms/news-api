package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewsArticleLang Модель перекладу статті.
type NewsArticleLang struct {
	Uuid             string `gorm:"primaryKey;type:uuid"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Rid              string         `gorm:"column:rid;type:string"`
	Loc              string         `gorm:"column:loc;type:string;size:5"`
	Title            string         `gorm:"column:title;type:string;size:255"`
	Content          string         `gorm:"column:content;type:string;size:64000"`
	ShortDescription string         `gorm:"column:short_description;type:string;size:1000"`
}

// BeforeCreate Генерація uuid.
func (n *NewsArticleLang) BeforeCreate(tx *gorm.DB) (err error) {
	n.Uuid = uuid.New().String()
	return
}
