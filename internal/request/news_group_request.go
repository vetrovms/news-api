package request

import (
	"news/internal/models"
	"news/internal/validator"
)

type NewsGroupRequest struct {
	Alias     string `json:"alias" validate:"required,max=255"`
	Published bool   `json:"published" validate:"omitempty,boolean"`
	Title     string `json:"title" validate:"required,max=255"`
}

func (r *NewsGroupRequest) Fill(m *models.NewsGroupDTO) {
	m.Title = r.Title
	m.Alias = r.Alias
	m.Published = r.Published // not required, default false
}

func (r *NewsGroupRequest) Validate() []string {
	return validator.Validate(r)
}
