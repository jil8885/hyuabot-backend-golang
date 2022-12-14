package cafeteria

import (
	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Menu{})
}
