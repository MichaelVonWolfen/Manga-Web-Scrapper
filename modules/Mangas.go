package modules

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	Name   string
	Url    string
	ImgUrl string

	Chapters []Chapter

	WebsiteID uint
	Website   Website `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
