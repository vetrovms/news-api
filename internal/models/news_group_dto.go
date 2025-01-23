package models

type NewsGroupDTO struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Alias     string          `json:"alias"`
	Published bool            `json:"published"`
	Files     []FileUploadDto `json:"files"`
}

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
