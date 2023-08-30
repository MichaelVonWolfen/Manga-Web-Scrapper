package helpers

import (
	"GoMangaWebScrapper/modules"
	"gorm.io/gorm"
)

func getWebsites() []modules.Website {
	var websites []modules.Website
	websites = append(websites, modules.Website{
		Model:         gorm.Model{},
		Name:          "Asura Scans",
		Url:           "https://asuracomics.com/",
		HasSearchPage: true,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "Manganato",
		Url:           "https://chapmanganato.com/",
		HasSearchPage: true,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "ReaperScans",
		Url:           "https://reaperscans.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "zeroscans",
		Url:           "https://zeroscans.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "darkscans",
		Url:           "https://darkscans.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "flamescans",
		Url:           "https://flamescans.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "toonily.net",
		Url:           "https://toonily.net/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "manhwaworld",
		Url:           "https://manhwaworld.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "https://1stkissmanga.me/",
		Url:           "https://1stkissmanga.me//",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "webtoons",
		Url:           "https://webtoons.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "zinmanga",
		Url:           "https://zinmanga.com/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "mm-scans.org",
		Url:           "https://mm-scans.org/",
		HasSearchPage: false,
	}, modules.Website{
		Model:         gorm.Model{},
		Name:          "manhuaplus",
		Url:           "https://manhuaplus.com/",
		HasSearchPage: false,
	})
	return websites
}
func getWebsitePatterns(tx *gorm.DB) []modules.WebsitePattern {
	websites := modules.Website{}
	tx.Find(&websites)

}
