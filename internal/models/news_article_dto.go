package models

type NewsArticleDTO struct {
	ID               int             `json:"id"`
	Alias            string          `json:"alias"`
	Published        bool            `json:"published"`
	Title            string          `json:"title"`
	Content          string          `json:"content"`
	ShortDescription string          `json:"short_description"`
	PublishedAt      string          `json:"published_at"`
	GroupId          int             `json:"group_id"`
	Group            NewsGroupDTO    `json:"group"`
	Files            []FileUploadDto `json:"files"`
}

func (dto *NewsArticleDTO) FillModel(model *NewsArticle, locale string) {
	model.Alias = dto.Alias
	model.Published = dto.Published
	model.PublishedAt = dto.PublishedAt
	model.GroupId = dto.GroupId
	if model.DefaultTitle == "" {
		model.DefaultTitle = dto.Title
	}
	if model.CurLang.Loc == "" {
		model.CurLang.Loc = locale
	}
	if model.CurLang.Loc == locale {
		model.CurLang.Title = dto.Title
		model.CurLang.Content = dto.Content
		model.CurLang.ShortDescription = dto.ShortDescription
	}
}
