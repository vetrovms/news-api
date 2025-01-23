package services

import (
	"context"
	"errors"
	"news/internal/database/repository"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
)

// NewsArticleService Сервіс новин.
type NewsArticleService struct {
	repo *repository.Repo
}

// NewNewsArticleService Конструктор сервіса новин.
func NewNewsArticleService(repo *repository.Repo) NewsArticleService {
	return NewsArticleService{
		repo: repo,
	}
}

// List Повертає список новин.
func (s *NewsArticleService) List(ctx context.Context, params map[string]string, locale string) (*[]models.NewsArticleDTO, error) {
	articles, err := s.repo.NewsArticleList(ctx, params, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := make([]models.NewsArticleDTO, len(articles))
	for i, m := range articles {
		dto[i] = m.DTO()
	}
	return &dto, nil
}

// One Повертає новину за ідентифікатором.
func (s *NewsArticleService) One(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOne(ctx, id, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := model.DTO()
	return &dto, nil
}

// Exists Перевіряє існування запису за ідентифікатором.
func (s *NewsArticleService) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := s.repo.NewsArticleExists(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// ExistsUnscoped Перевіряє існування м'яко видаленого запису за ідентифікатором.
func (s *NewsArticleService) ExistsUnscoped(ctx context.Context, id int) (bool, error) {
	exists, err := s.repo.NewsArticleExistsUnscoped(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// Create Створює нову новину.
func (s *NewsArticleService) Create(ctx context.Context, dto models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error) {
	var model models.NewsArticle
	dto.FillModel(&model, locale)
	err := s.repo.NewsArticleSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto = model.DTO()
	return &dto, nil
}

// Update Оновлює новину.
func (s *NewsArticleService) Update(ctx context.Context, dto models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOne(ctx, dto.ID, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto.FillModel(&model, locale)
	err = s.repo.NewsArticleSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto = model.DTO()
	return &dto, nil
}

// Trash М'яке видалення новини.
func (s *NewsArticleService) Trash(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOne(ctx, dto.ID, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsArticleTrash(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}

// Recover Відновлення новини після м'якого видалення.
func (s *NewsArticleService) Recover(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOneUnscoped(ctx, dto.ID, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsArticleRecover(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}

// Delete Остаточне видалення новини.
func (s *NewsArticleService) Delete(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOneUnscoped(ctx, dto.ID, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsArticleDelete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}
