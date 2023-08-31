package main

import (
	"GoMangaWebScrapper/helpers"
	"GoMangaWebScrapper/modules"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	db := modules.GetDatabaseInstance()

	helpers.InsertWebsitesAndQueries(db)

	//manga := modules.Manga{
	//	Model:   gorm.Model{},
	//	Name:    "Test2",
	//	Url:     "none",
	//	ImgUrl:  "example.com",
	//	Website: website,
	//}
	//chapter := modules.Chapter{
	//	Model:       gorm.Model{},
	//	Url:         "TestChapter",
	//	WasRead:     false,
	//	ChapterName: "Chapter 1 -  A new Saga for the BEBICI",
	//	Manga:       manga,
	//}
	//if err := tx.Create(&manga).Error; err != nil {
	//	// Rollback the transaction if there's an error
	//	tx.Rollback()
	//	fmt.Println("Error creating manga:", err)
	//	return
	//}
	//if err := tx.Create(&chapter).Error; err != nil {
	//	// Rollback the transaction if there's an error
	//	tx.Rollback()
	//	fmt.Println("Error creating manga:", err)
	//	return
	//}
	fmt.Println("Hello World")
	// Commit the transaction if everything is successful

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})
	// defining a data structure to store the scraped data
	c.OnHTML(".single-page > div > ul > li", func(e *colly.HTMLElement) {
		fmt.Println("e.Name")
		fmt.Println(e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	// downloading the target HTML page
	c.Visit("https://1stkissmanga.me/manga/the-ladys-law-of-survival/")

	/**
	Site-uri si statusul de accept al acestora:
		Reaperscan		NU
		ManhwaWorld		NU
		zinmanga		NU
		AsuraScans		DA
		chapmanganato	DA
		Webtoons		DA
		flamescans		DA
		zeroscans		DA
		darkscans		DA
		1stkissmanga	DA

	*/

}
