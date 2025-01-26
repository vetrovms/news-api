package services

import (
	"context"
	"errors"
	"news/internal/database/repository"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"news/internal/request"

	"github.com/gofiber/fiber/v2"
)

// NewsGroupService Сервіс груп новин.
type NewsGroupService struct {
	repo repository.IRepo
}

// NewNewsGroupService Конструктор сервіса груп новин.
func NewNewsGroupService(repo repository.IRepo) NewsGroupService {
	return NewsGroupService{
		repo: repo,
	}
}

// List Повертає список груп новин.
func (s *NewsGroupService) List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsGroupDTO, error) {
	groups, err := s.repo.NewsGroupList(ctx, c.Queries(), locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	groupsDto := make([]models.NewsGroupDTO, len(groups))
	for i, g := range groups {
		groupsDto[i] = g.DTO()
	}
	return &groupsDto, nil
}

// One Повертає групу новин за ідентифікатором.
func (s *NewsGroupService) One(ctx context.Context, c *fiber.Ctx, id int) (*models.NewsGroupDTO, error) {
	group, err := s.repo.NewsGroupOne(ctx, id, locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto := group.DTO()
	return &dto, nil
}

// Exists Перевіряє існування запису за ідентифікатором.
func (s *NewsGroupService) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := s.repo.NewsGroupExists(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// ExistsUnscoped Перевіряє існування м'яко видаленого запису за ідентифікатором.
func (s *NewsGroupService) ExistsUnscoped(ctx context.Context, id int) (bool, error) {
	exists, err := s.repo.NewsGroupExistsUnscoped(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New(myerrors.ServiceNotAvailable)
	}
	return exists, nil
}

// Create Створює нову групу новин.
func (s *NewsGroupService) Create(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest) (*models.NewsGroupDTO, error) {
	var model models.NewsGroup
	var dto models.NewsGroupDTO
	req.Fill(&dto)

	dto.FillModel(&model, locale(c))
	err := s.repo.NewsGroupSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto = model.DTO()
	return &dto, nil
}

// Update Оновлює групу новин.
func (s *NewsGroupService) Update(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest, id int) (*models.NewsGroupDTO, error) {
	var dto models.NewsGroupDTO
	req.Fill(&dto)
	dto.ID = id
	model, err := s.repo.NewsGroupOne(ctx, dto.ID, locale(c))
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto.FillModel(&model, locale(c))
	err = s.repo.NewsGroupSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto = model.DTO()
	return &dto, nil
}

// Trash М'яке видалення групи новин.
func (s *NewsGroupService) Trash(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOne(ctx, id, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupTrash(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}

// Recover Відновлення групи новин після м'якого видалення.
func (s *NewsGroupService) Recover(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOneUnscoped(ctx, id, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupRecover(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}

// Delete Остаточне видалення групи новин.
func (s *NewsGroupService) Delete(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOneUnscoped(ctx, id, locale)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupDelete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	r := model.DTO()
	return &r, nil
}
