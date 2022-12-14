package subway

type Timetable struct {
	StationID         string       `gorm:"column:station_id;primaryKey"`
	TerminalStationID string       `gorm:"column:terminal_station_id"`
	TerminalStation   RouteStation `gorm:"foreignKey:StationID;references:TerminalStationID"`
	DepartureTime     string       `gorm:"column:departure_time"`
	Weekday           string       `gorm:"column:weekday"`
	Heading           string       `gorm:"column:up_down_type"`
}

func (Timetable) TableName() string {
	return "subway_timetable"
}
