package models

// NewsGroupDTO DTO групи новин.
type NewsGroupDTO struct {
	ID        int             `json:"id" example:"321"`
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
