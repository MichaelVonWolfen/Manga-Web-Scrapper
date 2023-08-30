package helpers

import (
	"GoMangaWebScrapper/modules"
	"fmt"
	"gorm.io/gorm"
)

func InsertWebsitesAndQueries(db *gorm.DB) {
	tx := db.Begin()

	website := modules.Website{}
	//chapter := modules.Chapter{}
	//websitePattern := modules.WebsitePattern{}
	// Check if ther is already data in the tables and Rollback the transaction if there's an error
	if result := tx.First(&website); result.RowsAffected == 0 {
		websites := getWebsites()

		if err := tx.Create(&websites).Error; err != nil {
			tx.Rollback()
			panic(fmt.Sprintf("Error saving websites:\n %v", err))
			return
		}
	}
	fmt.Println("Hello World")
	// Commit the transaction if everything is successful
	tx.Commit()
}
