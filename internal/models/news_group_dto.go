package models

type NewsGroupDTO struct {
	ID        int              `json:"id"`
	Title     string           `json:"title"`
	Alias     string           `json:"alias"`
	Published bool             `json:"published"`
	Files     []*FileUploadDto `json:"files"`
}

func (dto *NewsGroupDTO) FillModel(g *NewsGroup, loc string) {
	g.Alias = dto.Alias
	g.Published = dto.Published
	if g.DefaultTitle == "" {
		g.DefaultTitle = dto.Title
	}
	if g.CurLang.Loc == "" {
		g.CurLang.Loc = loc
	}
	if g.CurLang.Loc == loc {
		g.CurLang.Title = dto.Title
	}
}
