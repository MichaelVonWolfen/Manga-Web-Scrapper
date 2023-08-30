package modules

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	Name            string
	Url             string
	HasSearchPage   bool
	Mangas          []Manga
	WebsitePatterns []WebsitePattern
}
