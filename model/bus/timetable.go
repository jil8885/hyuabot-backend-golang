package bus

type Timetable struct {
	RouteID       int    `gorm:"column:route_id;primaryKey"`
	StartStopID   int    `gorm:"column:start_stop_id;primaryKey"`
	Weekday       string `gorm:"column:weekday;primaryKey"`
	DepartureTime string `gorm:"column:departure_time;primaryKey"`
}

func (Timetable) TableName() string {
	return "bus_timetable"
}
