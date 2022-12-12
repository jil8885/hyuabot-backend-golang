package shuttle

import "gorm.io/gorm"

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&Route{})
	db.AutoMigrate(&Stop{})
	db.AutoMigrate(&Period{})
	db.AutoMigrate(&RouteStop{})
	db.AutoMigrate(&Timetable{})
}
