package services

import (
	"context"
	"errors"
	"fmt"
	"news/internal/config"
	"news/internal/database/repository"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"news/internal/request"
	"os"
	"slices"

	"github.com/gofiber/fiber/v2"
)

// FileUploadService Сервіс файлів.
type FileUploadService struct {
	repo repository.IRepo
}

// NewFileUploadService Конструктор сервіса файлів.
func NewFileUploadService(repo repository.IRepo) FileUploadService {
	return FileUploadService{
		repo: repo,
	}
}

// List Список файлів.
func (s *FileUploadService) List(ctx context.Context) (*[]models.FileUploadDto, error) {
	files, err := s.repo.FileUploadList(ctx)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := make([]models.FileUploadDto, len(files))
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
	dto := model.DTO()
	return &dto, nil
}

// Create Зберігає файл.
func (s *FileUploadService) Create(ctx context.Context, c *fiber.Ctx, req request.FileUploadRequest) (*models.FileUploadDto, error, int) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Log().Warn(err)
		return nil, err, fiber.StatusInternalServerError
	}

	allowedTypes := []string{"image/jpeg", "image/png"}
	if !slices.Contains(allowedTypes, file.Header.Get("Content-Type")) {
		logger.Log().Warn(myerrors.WrongFileFormat)
		return nil, errors.New(myerrors.WrongFileFormat), fiber.StatusBadRequest
	}

	destination := fmt.Sprintf(config.NewEnv().UploadPath+"%s", file.Filename)
	if err := c.SaveFile(file, destination); err != nil {
		logger.Log().Warn(err)
		return nil, err, fiber.StatusInternalServerError
	}

	var model models.FileUpload
	var dto models.FileUploadDto
	req.File = file
	req.Fill(&dto)
	dto.FillModel(&model)
	err = s.repo.FileUploadSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable), fiber.StatusInternalServerError
	}
	dto = model.DTO()
	return &dto, nil, fiber.StatusOK
}

// Delete Видалення файла.
func (s *FileUploadService) Delete(ctx context.Context, id int) (*models.FileUploadDto, error) {
	model, err := s.repo.FileUploadOne(ctx, id)
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
	r := model.DTO()
	return &r, nil
}
