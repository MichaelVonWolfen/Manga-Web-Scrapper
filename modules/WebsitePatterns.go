package modules

import "gorm.io/gorm"

type WebsitePattern struct {
	gorm.Model
	SearchQuery string
	WebsiteID   uint
	Website     Website `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
