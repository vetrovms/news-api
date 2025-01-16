package repository

import (
	"context"
	"news/internal/models"

	"gorm.io/gorm"
)

// NewsGroupList Повертає список груп новин.
func (repo *Repo) NewsGroupList(ctx context.Context, params map[string]string, loc string) ([]models.NewsGroup, error) {
	// repo.db.WithContext(ctx).Exec("select pg_sleep(10);") // @debug
	groups := []models.NewsGroup{}
	err := repo.db.WithContext(ctx).Preload("CurLang", "loc = ?", loc).Find(&groups).Error
	return groups, err
}

// NewsGroupExists Перевірка існування групи за ідентифікатором.
func (repo *Repo) NewsGroupExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Model(models.NewsGroup{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

// NewsGroupExists Перевірка існування м'яко видаленої групи за ідентифікатором.
func (repo *Repo) NewsGroupExistsUnscoped(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := repo.db.WithContext(ctx).Unscoped().Model(models.NewsGroup{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

// NewsGroupOne Повертає групу новин за ідентифікатором.
func (repo *Repo) NewsGroupOne(ctx context.Context, id int, loc string) (models.NewsGroup, error) {
	group := models.NewsGroup{}
	err := repo.db.WithContext(ctx).Preload("CurLang", "loc = ?", loc).First(&group, id).Error
	return group, err
}

// NewsGroupOneUnscoped Повертає м'яко видалену групу новин за ідентифікатором.
func (repo *Repo) NewsGroupOneUnscoped(ctx context.Context, id int, loc string) (models.NewsGroup, error) {
	group := models.NewsGroup{}
	err := repo.db.WithContext(ctx).Unscoped().Preload("CurLang", "loc = ?", loc).First(&group, id).Error
	return group, err
}

// NewsGroupSave Зберігає групу новин.
func (repo *Repo) NewsGroupSave(ctx context.Context, g *models.NewsGroup) error {
	return repo.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(&g).Error
}

// NewsGroupTrash М'яке видалення групи новин.
func (repo *Repo) NewsGroupTrash(ctx context.Context, g *models.NewsGroup) error {
	return repo.db.WithContext(ctx).Delete(&g).Error
}

// NewsGroupRecover Відновлення групи новин після м'якого видалення.
func (repo *Repo) NewsGroupRecover(ctx context.Context, g *models.NewsGroup) error {
	return repo.db.WithContext(ctx).Unscoped().Model(&g).Update("DeletedAt", nil).Error
}

// NewsGroupDelete Остаточне видалення групи новин.
func (repo *Repo) NewsGroupDelete(ctx context.Context, g *models.NewsGroup) error {
	return repo.db.WithContext(ctx).Unscoped().Delete(&g).Error
}
