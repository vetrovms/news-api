package repository

import (
	"gorm.io/gorm"
)

// Repo Репозиторій.
type Repo struct {
	db *gorm.DB
}

// NewRepo Конструктор репозиторія.
func NewRepo(conn *gorm.DB) Repo {
	return Repo{
		db: conn,
	}
}
