package models

// FileUploadDto DTO для моделі файла.
type FileUploadDto struct {
	ID         int    `json:"id"`
	EntityType string `json:"entity_type"`
	EntityId   int    `json:"entity_id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
}

// FillModel Заповнює поля моделі.
func (dto *FileUploadDto) FillModel(model *FileUpload) {
	model.EntityType = dto.EntityType
	model.EntityId = dto.EntityId
	model.Name = dto.Name
	model.Path = dto.Path
}
