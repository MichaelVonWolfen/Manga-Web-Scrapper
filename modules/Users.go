package modules

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}
