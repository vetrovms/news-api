package services

import (
	"context"
	"errors"
	"news/internal/config"
	"news/internal/database/repository"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"os"
)

// FileUploadService Сервіс файлів.
type FileUploadService struct {
	repo *repository.Repo
}

// NewFileUploadService Конструктор сервіса файлів.
func NewFileUploadService(repo *repository.Repo) FileUploadService {
	return FileUploadService{
		repo: repo,
	}
}

// List Список файлів.
func (s *FileUploadService) List(ctx context.Context) (*[]*models.FileUploadDto, error) {
	files, err := s.repo.FileUploadList(ctx)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	// dto := []*models.FileUploadDto{}
	dto := make([]*models.FileUploadDto, len(files))
	for i, m := range files {
		dto[i] = m.DTO()
	}
	return &dto, nil
}

// Exists Перевіряє існування запису за ідентифікатором.
func (s *FileUploadService) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := s.repo.FileUploadExists(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// One Повертає файл за ідентифікатором.
func (s *FileUploadService) One(ctx context.Context, id int) (*models.FileUploadDto, error) {
	model, err := s.repo.FileUploadOne(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Create Зберігає файл.
func (s *FileUploadService) Create(ctx context.Context, dto models.FileUploadDto) (*models.FileUploadDto, error) {
	var model models.FileUpload
	dto.FillModel(&model)
	err := s.repo.FileUploadSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Delete Видалення файла.
func (s *FileUploadService) Delete(ctx context.Context, dto *models.FileUploadDto) (*models.FileUploadDto, error) {
	model, err := s.repo.FileUploadOne(ctx, dto.ID)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.FileUploadDelete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = os.Remove(config.NewEnv().UploadPath + model.Name)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}
