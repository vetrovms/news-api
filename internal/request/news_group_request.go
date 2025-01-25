package request

import (
	"news/internal/models"
	"news/internal/validator"
)

// NewsGroupRequest Запит групи новин.
type NewsGroupRequest struct {
	Alias     string `json:"alias" validate:"required,max=255" example:"politics_ukraine"`
	Published bool   `json:"published" validate:"omitempty,boolean" example:"true"`
	Title     string `json:"title" validate:"required,max=255" example:"Політика, Україна"`
}

// Fill Заповнює DTO.
func (r *NewsGroupRequest) Fill(m *models.NewsGroupDTO) {
	m.Title = r.Title
	m.Alias = r.Alias
	m.Published = r.Published
}

// Validate Валідує запит групи новин.
func (r *NewsGroupRequest) Validate() []string {
	return validator.Validate(r)
}
