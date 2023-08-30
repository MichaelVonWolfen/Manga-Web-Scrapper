package modules

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Url         string
	WasRead     bool
	ChapterName string

	MangaID uint
	Manga   Manga `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
