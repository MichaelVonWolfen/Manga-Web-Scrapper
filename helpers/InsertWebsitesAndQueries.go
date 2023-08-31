package helpers

import (
	"gorm.io/gorm"
)

func InsertWebsitesAndQueries(db *gorm.DB) {
	tx := db.Begin()

	//chapter := modules.Chapter{}
	//websitePattern := modules.WebsitePattern{}
	seedWebsitesSeeds(tx)
	seedWebsitePatterns(tx)
	seedMangas(tx)
	// Commit the transaction if everything is successful
	tx.Commit()
	panic("FORCE STOP")
}
