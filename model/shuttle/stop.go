package shuttle

type Stop struct {
	Name      string      `gorm:"column:stop_name;primaryKey"`
	Latitude  float64     `gorm:"column:latitude"`
	Longitude float64     `gorm:"column:longitude"`
	RouteList []RouteStop `gorm:"foreignKey:StopName;references:Name"`
}

type StopItem struct {
	Name      string  `gorm:"column:stop_name;primaryKey"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
}

func (Stop) TableName() string {
	return "shuttle_stop"
}
