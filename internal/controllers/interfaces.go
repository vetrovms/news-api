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
	One(ctx context.Context, c *fiber.Ctx, uuid string) (*models.NewsArticleDTO, error)
	OneUnscoped(ctx context.Context, c *fiber.Ctx, uuid string) (*models.NewsArticleDTO, error)
	Exists(ctx context.Context, uuid string) (bool, error)
	ExistsUnscoped(ctx context.Context, uuid string) (bool, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest) (*models.NewsArticleDTO, error)
	Update(ctx context.Context, c *fiber.Ctx, req request.NewsArticleRequest, uuid string) (*models.NewsArticleDTO, error)
	Trash(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error)
	Recover(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error)
	Delete(ctx context.Context, uuid string, locale string) (*models.NewsArticleDTO, error)
}

// GroupsService інтерфейс сервіса груп новин.
type GroupsService interface {
	List(ctx context.Context, c *fiber.Ctx) (*[]models.NewsGroupDTO, error)
	One(ctx context.Context, c *fiber.Ctx, uuid string) (*models.NewsGroupDTO, error)
	Exists(ctx context.Context, uuid string) (bool, error)
	ExistsUnscoped(ctx context.Context, uuid string) (bool, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest) (*models.NewsGroupDTO, error)
	Update(ctx context.Context, c *fiber.Ctx, req request.NewsGroupRequest, uuid string) (*models.NewsGroupDTO, error)
	Trash(ctx context.Context, uuid string, locale string) (*models.NewsGroupDTO, error)
	Recover(ctx context.Context, uuid string, locale string) (*models.NewsGroupDTO, error)
	Delete(ctx context.Context, uuid string, locale string) (*models.NewsGroupDTO, error)
}

// FilesService Інтерфейс сервіса завантаження файлів.
type FilesService interface {
	List(ctx context.Context) (*[]models.FileUploadDto, error)
	Exists(ctx context.Context, uuid string) (bool, error)
	One(ctx context.Context, uuid string) (*models.FileUploadDto, error)
	Create(ctx context.Context, c *fiber.Ctx, req request.FileUploadRequest) (*models.FileUploadDto, error, int)
	Delete(ctx context.Context, uuid string) (*models.FileUploadDto, error)
}
