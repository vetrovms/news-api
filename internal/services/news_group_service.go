package services

import (
	"context"
	"errors"
	"news/internal/database/repository"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
)

// NewsGroupService Сервіс груп новин.
type NewsGroupService struct {
	repo *repository.Repo
}

// NewNewsGroupService Конструктор сервіса груп новин.
func NewNewsGroupService(repo *repository.Repo) NewsGroupService {
	return NewsGroupService{
		repo: repo,
	}
}

// List Повертає список груп новин.
func (s *NewsGroupService) List(ctx context.Context, params map[string]string, loc string) (*[]*models.NewsGroupDTO, error) {
	groups, err := s.repo.NewsGroupList(ctx, params, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	groupsDto := []*models.NewsGroupDTO{}
	for _, g := range groups {
		groupsDto = append(groupsDto, g.DTO())
	}
	return &groupsDto, nil
}

// One Повертає групу новин за ідентифікатором.
func (s *NewsGroupService) One(ctx context.Context, id int, loc string) (*models.NewsGroupDTO, error) {
	group, err := s.repo.NewsGroupOne(ctx, id, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return group.DTO(), nil
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
func (s *NewsGroupService) Create(ctx context.Context, dto models.NewsGroupDTO, loc string) (*models.NewsGroupDTO, error) {
	var model models.NewsGroup
	dto.FillModel(&model, loc)
	err := s.repo.NewsGroupSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Update Оновлює групу новин.
func (s *NewsGroupService) Update(ctx context.Context, dto models.NewsGroupDTO, loc string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOne(ctx, dto.ID, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	dto.FillModel(&model, loc)
	err = s.repo.NewsGroupSave(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Trash М'яке видалення групи новин.
func (s *NewsGroupService) Trash(ctx context.Context, dto *models.NewsGroupDTO, loc string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOne(ctx, dto.ID, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupTrash(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Recover Відновлення групи новин після м'якого видалення.
func (s *NewsGroupService) Recover(ctx context.Context, dto *models.NewsGroupDTO, loc string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOneUnscoped(ctx, dto.ID, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupRecover(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}

// Delete Остаточне видалення групи новин.
func (s *NewsGroupService) Delete(ctx context.Context, dto *models.NewsGroupDTO, loc string) (*models.NewsGroupDTO, error) {
	model, err := s.repo.NewsGroupOneUnscoped(ctx, dto.ID, loc)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	err = s.repo.NewsGroupDelete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New(myerrors.ServiceNotAvailable)
	}
	return model.DTO(), nil
}
