package models

import (
	"gorm.io/gorm"
)

type NewsGroup struct {
	gorm.Model
	Alias        string          `gorm:"column:alias;type:string;size:255"`
	Published    bool            `gorm:"column:published;type:bool"`
	DefaultTitle string          `gorm:"column:default_title;type:string;size:255"`
	Langs        []NewsGroupLang `gorm:"foreignKey:rid;references:id"`
	CurLang      NewsGroupLang   `gorm:"foreignKey:rid;references:id"`
	Files        []FileUpload    `gorm:"polymorphicType:EntityType;polymorphicId:EntityId;polymorphicValue:news_groups"`
}

func (g *NewsGroup) DTO() *NewsGroupDTO {
	filesDTO := make([]*FileUploadDto, len(g.Files))
	for i, f := range g.Files {
		filesDTO[i] = f.DTO()
	}
	return &NewsGroupDTO{
		ID:        int(g.ID),
		Title:     g.Title(),
		Alias:     g.Alias,
		Published: g.Published,
		Files:     filesDTO,
	}
}

func (g *NewsGroup) Title() string {
	if g.CurLang.Title != "" {
		return g.CurLang.Title
	}
	return g.DefaultTitle
}
