package repository

import (
	"context"
	"news/internal/models"
)

// FileUploadSave Зберігає запис про файл.
func (repo *Repo) FileUploadSave(ctx context.Context, model *models.FileUpload) error {
	return repo.db.WithContext(ctx).Save(&model).Error
}

// FileUploadList Список файлів.
func (repo *Repo) FileUploadList(ctx context.Context) ([]models.FileUpload, error) {
	models := []models.FileUpload{}
	err := repo.db.WithContext(ctx).Find(&models).Error
	return models, err
}

// FileUploadExists Перевірка існування файла за ідентифікатором.
func (repo *Repo) FileUploadExists(ctx context.Context, uuid string) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Model(models.FileUpload{}).Select("count(*) > 0").Where("uuid = ?", uuid).Find(&exists).Error
	return exists, err
}

// FileUploadOne Повертає файл за ідентифікатором.
func (repo *Repo) FileUploadOne(ctx context.Context, uuid string) (models.FileUpload, error) {
	model := models.FileUpload{}
	err := repo.db.WithContext(ctx).First(&model, "uuid = ?", uuid).Error
	return model, err
}

// FileUploadDelete Видалення файла.
func (repo *Repo) FileUploadDelete(ctx context.Context, model *models.FileUpload) error {
	return repo.db.WithContext(ctx).Unscoped().Delete(&model).Error
}
