package cafeteria

type Restaurant struct {
	CampusID     int     `gorm:"column:campus_id;primaryKey"`
	RestaurantID int     `gorm:"column:restaurant_id;primaryKey"`
	Name         string  `gorm:"column:restaurant_name"`
	Latitude     float64 `gorm:"column:latitude"`
	Longitude    float64 `gorm:"column:longitude"`
	MenuList     []Menu  `gorm:"foreignKey:RestaurantID;references:RestaurantID"`
}

type RestaurantItem struct {
	RestaurantID int    `gorm:"column:restaurant_id;primaryKey"`
	Name         string `gorm:"column:restaurant_name"`
	MenuList     []Menu `gorm:"foreignKey:RestaurantID;references:RestaurantID"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}
