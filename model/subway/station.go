package subway

type RouteStation struct {
	StationID       string      `gorm:"column:station_id;primaryKey"`
	RouteID         int         `gorm:"column:route_id;primaryKey"`
	StationName     string      `gorm:"column:station_name"`
	StationSequence int         `gorm:"column:station_sequence"`
	CumulativeTime  int         `gorm:"column:cumulative_time"`
	RealtimeList    []Realtime  `gorm:"foreignKey:StationID;references:StationID"`
	TimetableList   []Timetable `gorm:"foreignKey:StationID;references:StationID"`
}

type RouteStationListItem struct {
	StationID   string `gorm:"column:station_id;primaryKey"`
	RouteID     int    `gorm:"column:route_id;primaryKey"`
	StationName string `gorm:"column:station_name"`
}

type RouteStationItem struct {
	StationID     string      `gorm:"column:station_id;primaryKey"`
	RouteID       int         `gorm:"column:route_id;primaryKey"`
	StationName   string      `gorm:"column:station_name"`
	RealtimeList  []Realtime  `gorm:"foreignKey:StationID;references:StationID"`
	TimetableList []Timetable `gorm:"foreignKey:StationID;references:StationID"`
}

func (RouteStation) TableName() string {
	return "subway_route_station"
}
