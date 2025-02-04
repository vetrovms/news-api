package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewsGroup Модель групи новин.
type NewsGroup struct {
	Uuid         string `gorm:"primaryKey;type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt  `gorm:"index"`
	Alias        string          `gorm:"column:alias;type:string;size:255"`
	Published    bool            `gorm:"column:published;type:bool"`
	DefaultTitle string          `gorm:"column:default_title;type:string;size:255"`
	Langs        []NewsGroupLang `gorm:"foreignKey:rid;references:uuid"`
	CurLang      NewsGroupLang   `gorm:"foreignKey:rid;references:uuid"`
	Files        []FileUpload    `gorm:"polymorphicType:EntityType;polymorphicId:EntityId;polymorphicValue:news_groups"`
}

// DTO Повертає DTO групи новин.
func (g *NewsGroup) DTO() NewsGroupDTO {
	filesDTO := make([]FileUploadDto, len(g.Files))
	for i, f := range g.Files {
		filesDTO[i] = f.DTO()
	}
	return NewsGroupDTO{
		Uuid:      g.Uuid,
		Title:     g.Title(),
		Alias:     g.Alias,
		Published: g.Published,
		Files:     filesDTO,
	}
}

// Title Повертає заголовок групи новин на поточній мові.
func (g *NewsGroup) Title() string {
	if g.CurLang.Title != "" {
		return g.CurLang.Title
	}
	return g.DefaultTitle
}

// BeforeCreate Генерація uuid.
func (n *NewsGroup) BeforeCreate(tx *gorm.DB) (err error) {
	n.Uuid = uuid.New().String()
	return
}
