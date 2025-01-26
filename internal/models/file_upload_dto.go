package models

// FileUploadDto DTO для моделі файла.
type FileUploadDto struct {
	ID         int    `json:"id" example:"1"`
	EntityType string `json:"entity_type" example:"news_articles"`
	EntityId   int    `json:"entity_id" example:"123"`
	Name       string `json:"name" example:"article_img_123.png"`
	Path       string `json:"path" example:"/uploads/article_img_123.png"`
}

// FillModel Заповнює поля моделі.
func (dto *FileUploadDto) FillModel(model *FileUpload) {
	model.EntityType = dto.EntityType
	model.EntityId = dto.EntityId
	model.Name = dto.Name
	model.Path = dto.Path
}
