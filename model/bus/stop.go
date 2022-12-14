package bus

type Stop struct {
	StopID       int     `gorm:"column:stop_id;primaryKey"`
	StopName     string  `gorm:"column:stop_name"`
	DistrictCode int     `gorm:"column:district_code"`
	RegionName   string  `gorm:"column:region_name"`
	Latitude     float64 `gorm:"column:latitude"`
	Longitude    float64 `gorm:"column:longitude"`
}

func (Stop) TableName() string {
	return "bus_stop"
}
