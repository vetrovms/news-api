package models

import (
	"gorm.io/gorm"
)

// NewsGroupLand Модель перекладу групин новин.
type NewsGroupLang struct {
	gorm.Model
	Rid   int    `gorm:"column:rid;type:int"`
	Loc   string `gorm:"column:loc;type:string;size:5"`
	Title string `gorm:"column:title;type:string;size:255"`
}
