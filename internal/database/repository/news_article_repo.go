package repository

import (
	"context"
	"news/internal/models"

	"gorm.io/gorm"
)

// NewsArticleList Повертає список новин.
func (repo *Repo) NewsArticleList(ctx context.Context, params map[string]string, loc string) ([]models.NewsArticle, error) {
	// repo.db.WithContext(ctx).Exec("select pg_sleep(10);") // @debug
	models := []models.NewsArticle{}
	err := repo.db.WithContext(ctx).
		Preload("CurLang", "loc = ?", loc).
		Preload("Group.CurLang", "loc = ?", loc).
		Preload("Files").
		Find(&models).Error
	return models, err
}

// NewsArticleExists Перевірка існування новини за ідентифікатором.
func (repo *Repo) NewsArticleExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Model(models.NewsArticle{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

// NewsArticleExists Перевірка існування м'яко видаленої новини за ідентифікатором.
func (repo *Repo) NewsArticleExistsUnscoped(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Unscoped().Model(models.NewsArticle{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

// NewsArticleOne Повертає новину за ідентифікатором.
func (repo *Repo) NewsArticleOne(ctx context.Context, id int, loc string) (models.NewsArticle, error) {
	model := models.NewsArticle{}
	err := repo.db.WithContext(ctx).Preload("CurLang", "loc = ?", loc).Preload("Group.CurLang", "loc = ?", loc).Preload("Files").First(&model, id).Error
	return model, err
}

// NewsArticleOneUnscoped Повертає м'яко новину за ідентифікатором.
func (repo *Repo) NewsArticleOneUnscoped(ctx context.Context, id int, loc string) (models.NewsArticle, error) {
	model := models.NewsArticle{}
	err := repo.db.WithContext(ctx).Unscoped().
		Preload("CurLang", "loc = ?", loc).
		Preload("Group.CurLang", "loc = ?", loc).
		Preload("Files").First(&model, id).Error
	return model, err
}

// NewsArticleSave Зберігає новину.
func (repo *Repo) NewsArticleSave(ctx context.Context, model *models.NewsArticle) error {
	return repo.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(&model).Error
}

// NewsArticleTrash М'яке видалення новини.
func (repo *Repo) NewsArticleTrash(ctx context.Context, model *models.NewsArticle) error {
	return repo.db.WithContext(ctx).Delete(&model).Error
}

// NewsArticleRecover Відновлення новини після м'якого видалення.
func (repo *Repo) NewsArticleRecover(ctx context.Context, model *models.NewsArticle) error {
	return repo.db.WithContext(ctx).Unscoped().Model(&model).Update("DeletedAt", nil).Error
}

// NewsArticleDelete Остаточне видалення новини.
func (repo *Repo) NewsArticleDelete(ctx context.Context, model *models.NewsArticle) error {
	return repo.db.WithContext(ctx).Unscoped().Delete(&model).Error
}
