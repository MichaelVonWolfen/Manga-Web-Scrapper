package helpers

import (
	"GoMangaWebScrapper/modules"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func seedWebsitesSeeds(tx *gorm.DB) {
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
	if result := tx.First(&modules.Website{}); result.RowsAffected == 0 {
		if err := tx.Create(&websites).Error; err != nil {
			tx.Rollback()
			panic(fmt.Sprintf("Error saving websites:\n %v", err))
			return
		}
	} else {
		fmt.Println("There are already some Websites present")
	}
}
func seedWebsitePatterns(tx *gorm.DB) {
	var websites []modules.Website
	sitePatterns := make(map[string]string)
	sitePatterns["Asura Scans"] = "#chapterlist > ul"
	sitePatterns["Manganato"] = "div.panel-story-chapter-list > ul"
	sitePatterns["ReaperScans"] = "div.pb-4 > div > div > ul"
	sitePatterns["zeroscans"] = "div.row.pa-4"
	sitePatterns["darkscans"] = "#manga-chapters-holder >.single-page > div > ul"
	sitePatterns["flamescans"] = "#chapterlist > ul"
	sitePatterns["toonily.net"] = ".single-page > div > ul"
	sitePatterns["manhwaworld"] = ".single-page > div > ul"
	sitePatterns["https://1stkissmanga.me/"] = ".single-page > div > ul"
	sitePatterns["webtoons"] = "#_listUl"
	sitePatterns["zinmanga"] = ".single-page > div > ul"
	sitePatterns["mm-scans.org"] = ".chapter-content-list > div > ul"
	sitePatterns["manhuaplus"] = ".single-page > div > ul"

	tx.Find(&websites)
	var websitePatterns []modules.WebsitePattern
	for _, website := range websites {
		websitePatterns = append(websitePatterns, modules.WebsitePattern{
			Model:       gorm.Model{},
			SearchQuery: sitePatterns[website.Name],
			Website:     website,
		})
	}
	if result := tx.First(&modules.WebsitePattern{}); result.RowsAffected == 0 {
		if err := tx.Create(&websitePatterns).Error; err != nil {
			tx.Rollback()
			panic(fmt.Sprintf("Error saving websites patterns:\n %v", err))
		}
	} else {
		fmt.Println("There are already some Website Patterns present")
	}
}
func seedMangas(tx *gorm.DB) {
	var websites []modules.Website
	siteURLS := make(map[string]string)
	siteURLS["Asura Scans"] = "https://chapmanganato.com/manga-ci980191"
	siteURLS["Manganato"] = "https://reaperscans.com/comics/7319-the-great-mage-returns-after-4000-years"
	siteURLS["ReaperScans"] = "https://zeroscans.com/comics/second-life-ranker"
	siteURLS["zeroscans"] = "https://manhuaplus.com/manga/disastrous-necromancer"
	siteURLS["darkscans"] = "https://asuracomics.com/manga/9141111445-player-who-cant-level-up"
	siteURLS["flamescans"] = "https://zinmanga.com/manga/the-villains-precious-daughter"
	siteURLS["toonily.net"] = "https://www.webtoons.com/en/action/omniscient-reader/list?title_no=2154"
	siteURLS["manhwaworld"] = "https://1stkissmanga.me/manga/seduce-the-villains-father"
	siteURLS["https://1stkissmanga.me/"] = "https://toonily.net/manga/im-really-not-the-demon-gods-lackey"
	siteURLS["webtoons"] = "https://manhwaworld.com/manga/99-wooden-stick"
	siteURLS["zinmanga"] = "https://flamescans.org/series/1693389721-omniscient-readers-viewpoint	"
	siteURLS["mm-scans.org"] = "https://darkscans.com/mangas/player"
	siteURLS["manhuaplus"] = "https://zeroscans.com/comics/second-life-ranker"

	tx.Find(&websites)
	var mangas []modules.Manga
	for _, website := range websites {
		name := strings.Split(siteURLS[website.Name], "/")
		mangas = append(mangas, modules.Manga{
			Model:   gorm.Model{},
			Name:    name[len(name)-1],
			Url:     siteURLS[website.Name],
			Website: website,
		})
	}
	if result := tx.First(&modules.Manga{}); result.RowsAffected != 0 {
		fmt.Println("There are already some Mangas present")
	}
	if err := tx.Create(&mangas).Error; err != nil {
		tx.Rollback()
		panic(fmt.Sprintf("Error saving mangas:\n %v", err))
	}
}
