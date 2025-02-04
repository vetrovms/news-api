package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewsArticle Модель статті.
type NewsArticle struct {
	Uuid         string `gorm:"primaryKey;type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt    `gorm:"index"`
	Alias        string            `gorm:"column:alias;type:string;size:255"`
	Published    bool              `gorm:"column:published;type:bool"`
	DefaultTitle string            `gorm:"column:default_title;type:string;size:255"`
	PublishedAt  string            `gorm:"column:published_at;type:string"`
	GroupId      string            `gorm:"column:group_id;type:string"`
	UserId       int               `gorm:"column:user_id;type:int"`
	Group        NewsGroup         `gorm:"foreignKey:uuid;references:group_id"`
	Langs        []NewsArticleLang `gorm:"foreignKey:rid;references:uuid"`
	CurLang      NewsArticleLang   `gorm:"foreignKey:rid;references:uuid"`
	Files        []FileUpload      `gorm:"polymorphicType:EntityType;polymorphicId:EntityId;polymorphicValue:news_articles"`
}

// DTO Повертає DTO статті.
func (a *NewsArticle) DTO() NewsArticleDTO {
	filesDTO := make([]FileUploadDto, len(a.Files))
	for i, f := range a.Files {
		filesDTO[i] = f.DTO()
	}
	return NewsArticleDTO{
		Uuid:             a.Uuid,
		Title:            a.Title(),
		Content:          a.CurLang.Content,
		ShortDescription: a.CurLang.ShortDescription,
		Alias:            a.Alias,
		Published:        a.Published,
		PublishedAt:      a.PublishedAt,
		GroupId:          a.GroupId,
		UserId:           a.UserId,
		Group: NewsGroupDTO{
			Uuid:      a.GroupId,
			Title:     a.Group.Title(),
			Alias:     a.Group.Alias,
			Published: a.Group.Published,
		},
		Files: filesDTO,
	}
}

// Title Повертає заголовок статті на поточній мові.
func (a *NewsArticle) Title() string {
	if a.CurLang.Title != "" {
		return a.CurLang.Title
	}
	return a.DefaultTitle
}

// BeforeCreate Генерація uuid.
func (n *NewsArticle) BeforeCreate(tx *gorm.DB) (err error) {
	n.Uuid = uuid.New().String()
	return
}
