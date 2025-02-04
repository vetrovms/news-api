package services

import (
	"context"
	"errors"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"news/internal/request"

	"github.com/gofiber/fiber/v2"
)

// NewsArticleService Сервіс новин.
type NewsArticleService struct {
	repo newsRepo
}

// NewNewsArticleService Конструктор сервіса новин.
func NewNewsArticleService(repo newsRepo) NewsArticleService {
	return NewsArticleService{
		repo: repo,
	}
}

// List Повертає список новин.
func (s *NewsArticleService) List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsArticleDTO, error) {
	articles, err := s.repo.NewsArticleList(ctx, c.Queries(), locale(c))
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
func (s *NewsArticleService) One(ctx context.Context, c *fiber.Ctx, uuid string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOne(ctx, uuid, locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := model.DTO()
	return &dto, nil
}

// OneUnscoped Повертає видалену новину за ідентифікатором.
func (s *NewsArticleService) OneUnscoped(ctx context.Context, c *fiber.Ctx, uuid string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOneUnscoped(ctx, uuid, locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := model.DTO()
	return &dto, nil
}

// Exists Перевіряє існування запису за ідентифікатором.
func (s *NewsArticleService) Exists(ctx context.Context, uuid string) (bool, error) {
	exists, err := s.repo.NewsArticleExists(ctx, uuid)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// ExistsUnscoped Перевіряє існування м'яко видаленого запису за ідентифікатором.
func (s *NewsArticleService) ExistsUnscoped(ctx context.Context, uuid string) (bool, error) {
	exists, err := s.repo.NewsArticleExistsUnscoped(ctx, uuid)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// Create Створює нову статтю.
func (s *NewsArticleService) Create(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest) (*models.NewsArticleDTO, error) {
	var model models.NewsArticle
	var dto models.NewsArticleDTO

	jwtString := request.TokenFromRequest(c)
	claims, err := request.ClaimsFromToken(jwtString)

	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}

	userId := claims["sub"]
	dto.UserId = int(userId.(float64))
	req.Fill(&dto)
	dto.FillModel(&model, locale(c))

	err = s.repo.NewsArticleSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}

	dto = model.DTO()
	return &dto, nil
}

// Update Оновлює новину.
func (s *NewsArticleService) Update(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest, uuid string) (*models.NewsArticleDTO, error) {
	var dto models.NewsArticleDTO
	req.Fill(&dto)
	dto.Uuid = uuid

	model, err := s.repo.NewsArticleOne(ctx, dto.Uuid, locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}

	dto.FillModel(&model, locale(c))
	err = s.repo.NewsArticleSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}

	dto = model.DTO()
	return &dto, nil
}

// Trash М'яке видалення новини.
func (s *NewsArticleService) Trash(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOne(ctx, uuid, locale)
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
func (s *NewsArticleService) Recover(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOneUnscoped(ctx, uuid, locale)
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
func (s *NewsArticleService) Delete(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error) {
	model, err := s.repo.NewsArticleOneUnscoped(ctx, uuid, locale)
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
