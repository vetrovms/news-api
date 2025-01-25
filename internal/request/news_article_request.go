package request

import (
	"news/internal/models"
	"news/internal/validator"
)

// NewsArticleRequest Тіло запита статті новин.
type NewsArticleRequest struct {
	Alias            string `json:"alias" validate:"required,max=255" example:"nova_stattya"`
	Published        bool   `json:"published" validate:"omitempty,boolean" bool:"true"`
	Title            string `json:"title" validate:"required,max=255" example:"Нова новина"`
	Content          string `json:"content" validate:"required,max=64000" example:"Сьогодні щось відбулось."`
	ShortDescription string `json:"short_description" validate:"omitempty,max=1000" example:"Короткий опис новини."`
	PublishedAt      string `json:"published_at" form:"published_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z" example:"2006-01-02T15:04:05Z"`
	GroupId          int    `json:"group_id" form:"group_id" validate:"number" example:"3"`
}

// Fill Заповнює DTO.
func (r *NewsArticleRequest) Fill(dto *models.NewsArticleDTO) {
	dto.Title = r.Title
	dto.Alias = r.Alias
	dto.Published = r.Published
	dto.Content = r.Content
	dto.ShortDescription = r.ShortDescription
	dto.PublishedAt = r.PublishedAt
	dto.GroupId = r.GroupId
}

// Validate Валідує запит статті новин.
func (r *NewsArticleRequest) Validate() []string {
	return validator.Validate(r)
}
