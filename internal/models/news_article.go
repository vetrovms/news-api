package models

import (
	"gorm.io/gorm"
)

type NewsArticle struct {
	gorm.Model
	Alias        string            `gorm:"column:alias;type:string;size:255"`
	Published    bool              `gorm:"column:published;type:bool"`
	DefaultTitle string            `gorm:"column:default_title;type:string;size:255"`
	PublishedAt  string            `gorm:"column:published_at;type:string"`
	GroupId      int               `gorm:"column:group_id;type:int"`
	Group        NewsGroup         `gorm:"foreignKey:id;references:group_id"`
	Langs        []NewsArticleLang `gorm:"foreignKey:rid;references:id"`
	CurLang      NewsArticleLang   `gorm:"foreignKey:rid;references:id"`
	Files        []FileUpload      `gorm:"polymorphicType:EntityType;polymorphicId:EntityId;polymorphicValue:news_articles"`
}

func (a *NewsArticle) DTO() *NewsArticleDTO {
	filesDTO := make([]*FileUploadDto, len(a.Files))
	for i, f := range a.Files {
		filesDTO[i] = f.DTO()
	}
	return &NewsArticleDTO{
		ID:               int(a.ID),
		Title:            a.Title(),
		Content:          a.CurLang.Content,
		ShortDescription: a.CurLang.ShortDescription,
		Alias:            a.Alias,
		Published:        a.Published,
		PublishedAt:      a.PublishedAt,
		GroupId:          a.GroupId,
		Group: NewsGroupDTO{
			ID:        a.GroupId,
			Title:     a.Group.Title(),
			Alias:     a.Group.Alias,
			Published: a.Group.Published,
		},
		Files: filesDTO,
	}
}

func (a *NewsArticle) Title() string {
	if a.CurLang.Title != "" {
		return a.CurLang.Title
	}
	return a.DefaultTitle
}
