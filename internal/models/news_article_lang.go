package models

import (
	"gorm.io/gorm"
)

type NewsArticleLang struct {
	gorm.Model
	Rid              int    `gorm:"column:rid;type:int"`
	Loc              string `gorm:"column:loc;type:string;size:5"`
	Title            string `gorm:"column:title;type:string;size:255"`
	Content          string `gorm:"column:content;type:string;size:64000"`
	ShortDescription string `gorm:"column:short_description;type:string;size:1000"`
}
