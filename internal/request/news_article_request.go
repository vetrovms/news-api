package request

import (
	"news/internal/models"
	"news/internal/validator"
)

type NewsArticleRequest struct {
	Alias            string `json:"alias" validate:"required,max=255"`
	Published        bool   `json:"published" validate:"omitempty,boolean"`
	Title            string `json:"title" validate:"required,max=255"`
	Content          string `json:"content" validate:"required,max=64000"`
	ShortDescription string `json:"short_description" validate:"omitempty,max=1000"`
	PublishedAt      string `json:"published_at" form:"published_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z"`
	GroupId          int    `json:"group_id" form:"group_id" validate:"number"`
}

func (r *NewsArticleRequest) Fill(dto *models.NewsArticleDTO) {
	dto.Title = r.Title
	dto.Alias = r.Alias
	dto.Published = r.Published
	dto.Content = r.Content
	dto.ShortDescription = r.ShortDescription
	dto.PublishedAt = r.PublishedAt
	dto.GroupId = r.GroupId
}

func (r *NewsArticleRequest) Validate() []string {
	return validator.Validate(r)
}
