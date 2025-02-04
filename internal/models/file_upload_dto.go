package models

// FileUploadDto DTO для моделі файла.
type FileUploadDto struct {
	Uuid       string `json:"uuid" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
	EntityType string `json:"entity_type" example:"news_articles"`
	EntityId   string `json:"entity_id" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
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
