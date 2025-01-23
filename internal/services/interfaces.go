package services

import (
	"context"

	"news/internal/models"
)

type INewsService interface {
	List(ctx context.Context, params map[string]string, locale string) (*[]models.NewsArticleDTO, error)
	One(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
	Exists(ctx context.Context, id int) (bool, error)
	ExistsUnscoped(ctx context.Context, id int) (bool, error)
	Create(ctx context.Context, dto models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error)
	Update(ctx context.Context, dto models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error)
	Trash(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error)
	Recover(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error)
	Delete(ctx context.Context, dto *models.NewsArticleDTO, locale string) (*models.NewsArticleDTO, error)
}

type IGroupsService interface {
	List(ctx context.Context, params map[string]string, locale string) (*[]models.NewsGroupDTO, error)
	One(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error)
	Exists(ctx context.Context, id int) (bool, error)
	ExistsUnscoped(ctx context.Context, id int) (bool, error)
	Create(ctx context.Context, dto models.NewsGroupDTO, locale string) (*models.NewsGroupDTO, error)
	Update(ctx context.Context, dto models.NewsGroupDTO, locale string) (*models.NewsGroupDTO, error)
	Trash(ctx context.Context, dto *models.NewsGroupDTO, locale string) (*models.NewsGroupDTO, error)
	Recover(ctx context.Context, dto *models.NewsGroupDTO, locale string) (*models.NewsGroupDTO, error)
	Delete(ctx context.Context, dto *models.NewsGroupDTO, locale string) (*models.NewsGroupDTO, error)
}

type IFilesService interface {
	List(ctx context.Context) (*[]models.FileUploadDto, error)
	Exists(ctx context.Context, id int) (bool, error)
	One(ctx context.Context, id int) (*models.FileUploadDto, error)
	Create(ctx context.Context, dto models.FileUploadDto) (*models.FileUploadDto, error)
	Delete(ctx context.Context, dto *models.FileUploadDto) (*models.FileUploadDto, error)
}
