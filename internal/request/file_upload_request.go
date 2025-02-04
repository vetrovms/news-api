package request

import (
	"mime/multipart"
	"news/internal/config"
	"news/internal/models"
	"news/internal/validator"
)

// FileUploadRequest Тіло запита завантаження файла.
type FileUploadRequest struct {
	EntityType string                `json:"entity_type" form:"entity_type" validate:"oneof=news_groups news_articles" example:"news_article"`
	EntityId   string                `json:"entity_id" form:"entity_id" example:"0194cd77-d0ab-74db-88be-f9de341a4b5f"`
	File       *multipart.FileHeader `swaggerignore:"true"`
}

// Fill Заповнює DTO.
func (r *FileUploadRequest) Fill(dto *models.FileUploadDto) {
	dto.EntityId = r.EntityId
	dto.EntityType = r.EntityType
	dto.Name = r.File.Filename
	dto.Path = "/" + config.NewEnv().UploadDir + "/" + dto.Name
}

// Validate Валідує запит.
func (r *FileUploadRequest) Validate() []string {
	return validator.Validate(r)
}
