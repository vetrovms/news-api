package models

// NewsGroupDTO DTO групи новин.
type NewsGroupDTO struct {
	Uuid      string          `json:"uuid" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
	Title     string          `json:"title" example:"Спорт"`
	Alias     string          `json:"alias" example:"sport"`
	Published bool            `json:"published" example:"true"`
	Files     []FileUploadDto `json:"files"`
}

// FillModel Заповнює модель.
func (dto *NewsGroupDTO) FillModel(g *NewsGroup, locale string) {
	g.Alias = dto.Alias
	g.Published = dto.Published
	if g.DefaultTitle == "" {
		g.DefaultTitle = dto.Title
	}
	if g.CurLang.Loc == "" {
		g.CurLang.Loc = locale
	}
	if g.CurLang.Loc == locale {
		g.CurLang.Title = dto.Title
	}
}
