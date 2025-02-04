package models

// NewsArticleDTO DTO статті.
type NewsArticleDTO struct {
	Uuid             string          `json:"uuid" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
	Alias            string          `json:"alias" example:"new_article_uri"`
	Published        bool            `json:"published" example:"true"`
	Title            string          `json:"title" example:"Хороша новина"`
	Content          string          `json:"content" example:"Сталось щось добре"`
	ShortDescription string          `json:"short_description" example:"Короткий опис новини."`
	PublishedAt      string          `json:"published_at" example:"2024-12-05 12:48"`
	GroupId          string          `json:"group_id" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
	UserId           int             `json:"user_id" example:"456"`
	Group            NewsGroupDTO    `json:"group"`
	Files            []FileUploadDto `json:"files"`
}

// FillModel Заповнює модель.
func (dto *NewsArticleDTO) FillModel(model *NewsArticle, locale string) {
	model.Alias = dto.Alias
	model.Published = dto.Published
	model.PublishedAt = dto.PublishedAt
	model.GroupId = dto.GroupId
	if dto.UserId != 0 {
		model.UserId = dto.UserId
	}
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
