package services

import (
	"context"

	"news/internal/models"
)

type filesRepo interface {
	FileUploadSave(ctx context.Context, model *models.FileUpload) error
	FileUploadList(ctx context.Context) ([]models.FileUpload, error)
	FileUploadExists(ctx context.Context, uuid string) (bool, error)
	FileUploadOne(ctx context.Context, uuid string) (models.FileUpload, error)
	FileUploadDelete(ctx context.Context, model *models.FileUpload) error
}

type newsRepo interface {
	NewsArticleList(ctx context.Context, params map[string]string, locale string) ([]models.NewsArticle, error)
	NewsArticleExists(ctx context.Context, uuid string) (bool, error)
	NewsArticleExistsUnscoped(ctx context.Context, uuid string) (bool, error)
	NewsArticleOne(ctx context.Context, uuid string, locale string) (models.NewsArticle, error)
	NewsArticleOneUnscoped(ctx context.Context, uuid string, locale string) (models.NewsArticle, error)
	NewsArticleSave(ctx context.Context, model *models.NewsArticle) error
	NewsArticleTrash(ctx context.Context, model *models.NewsArticle) error
	NewsArticleRecover(ctx context.Context, model *models.NewsArticle) error
	NewsArticleDelete(ctx context.Context, model *models.NewsArticle) error
}

type groupsRepo interface {
	NewsGroupList(ctx context.Context, params map[string]string, locale string) ([]models.NewsGroup, error)
	NewsGroupExists(ctx context.Context, uuid string) (bool, error)
	NewsGroupExistsUnscoped(ctx context.Context, uuid string) (bool, error)
	NewsGroupOne(ctx context.Context, uuid string, locale string) (models.NewsGroup, error)
	NewsGroupOneUnscoped(ctx context.Context, uuid string, locale string) (models.NewsGroup, error)
	NewsGroupSave(ctx context.Context, g *models.NewsGroup) error
	NewsGroupTrash(ctx context.Context, g *models.NewsGroup) error
	NewsGroupRecover(ctx context.Context, g *models.NewsGroup) error
	NewsGroupDelete(ctx context.Context, g *models.NewsGroup) error
}
