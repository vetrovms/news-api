package controllers

import (
	"context"
	"news/internal/models"
	"news/internal/request"

	"github.com/gofiber/fiber/v2"
)

// NewsService Інтерфейс сервіса новин.
type NewsService interface {
	List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsArticleDTO, error)
	One(ctx context.Context, c *fiber.Ctx, id int) (*models.NewsArticleDTO, error)
	OneUnscoped(ctx context.Context, c *fiber.Ctx, id int) (*models.NewsArticleDTO, error)
	Exists(ctx context.Context, id int) (bool, error)
	ExistsUnscoped(ctx context.Context, id int) (bool, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest) (*models.NewsArticleDTO, error)
	Update(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest, id int) (*models.NewsArticleDTO, error)
	Trash(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
	Recover(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
	Delete(ctx context.Context, id int, locale string) (*models.NewsArticleDTO, error)
}

// GroupsService інтерфейс сервіса груп новин.
type GroupsService interface {
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

// FilesService Інтерфейс сервіса завантаження файлів.
type FilesService interface {
	List(ctx context.Context) (*[]models.FileUploadDto, error)
	Exists(ctx context.Context, id int) (bool, error)
	One(ctx context.Context, id int) (*models.FileUploadDto, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.FileUploadRequest) (*models.FileUploadDto, error, int)
	Delete(ctx context.Context, id int) (*models.FileUploadDto, error)
}
