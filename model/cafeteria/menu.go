package cafeteria

type Menu struct {
	RestaurantID int    `gorm:"column:restaurant_id;primaryKey"`
	FeedDate     string `gorm:"column:feed_date;primaryKey"`
	TimeType     string `gorm:"column:time_type;primaryKey"`
	Menu         string `gorm:"column:menu_food"`
	Price        string `gorm:"column:menu_price"`
}
