package services

import (
	"context"

	"news/internal/models"
	"news/internal/request"

	"github.com/gofiber/fiber/v2"
)

// INewsService Інтерфейс сервіса новин.
type INewsService interface {
	List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsArticleDTO, error)
	One(ctx context.Context, c *fiber.Ctx, id int) (*models.NewsArticleDTO, error)
	Exists(ctx context.Context, id int) (bool, error)
	ExistsUnscoped(ctx context.Context, id int) (bool, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest) (*models.NewsArticleDTO, error)
	Update(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest, id int) (*models.NewsArticleDTO, error)
	Trash(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
	Recover(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
	Delete(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
}

// IGroupsService інтерфейс сервіса груп новин.
type IGroupsService interface {
	List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsGroupDTO, error)
	One(ctx context.Context, c *fiber.Ctx, id int) (*models.NewsGroupDTO, error)
	Exists(ctx context.Context, id int) (bool, error)
	ExistsUnscoped(ctx context.Context, id int) (bool, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest) (*models.NewsGroupDTO, error)
	Update(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest, id int) (*models.NewsGroupDTO, error)
	Trash(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error)
	Recover(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error)
	Delete(ctx context.Context, id int, locale string) (*models.NewsGroupDTO, error)
}

// IFilesService Інтерфейс сервіса завантаження файлів.
type IFilesService interface {
	List(ctx context.Context) (*[]models.FileUploadDto, error)
	Exists(ctx context.Context, id int) (bool, error)
	One(ctx context.Context, id int) (*models.FileUploadDto, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.FileUploadRequest) (*models.FileUploadDto, error, int)
	Delete(ctx context.Context, id int) (*models.FileUploadDto, error)
}
