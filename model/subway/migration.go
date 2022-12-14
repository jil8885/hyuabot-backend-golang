package subway

import "gorm.io/gorm"

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&RouteStation{})
}
