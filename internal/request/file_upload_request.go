package request

import (
	"mime/multipart"
	"news/internal/config"
	"news/internal/models"
	"news/internal/validator"
)

type FileUploadRequest struct {
	EntityType string `form:"entity_type" validate:"oneof=news_groups news_articles"`
	EntityId   int    `form:"entity_id"`
	File       *multipart.FileHeader
}

func (r *FileUploadRequest) Fill(dto *models.FileUploadDto) {
	dto.EntityId = r.EntityId
	dto.EntityType = r.EntityType
	dto.Name = r.File.Filename
	dto.Path = "/" + config.NewEnv().UploadDir + "/" + dto.Name
}

func (r *FileUploadRequest) Validate() []string {
	return validator.Validate(r)
}
