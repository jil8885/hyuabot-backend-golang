package subway

type RouteStation struct {
	StationID       string `gorm:"column:station_id;primaryKey"`
	RouteID         int    `gorm:"column:route_id;primaryKey"`
	StationName     string `gorm:"column:station_name"`
	StationSequence int    `gorm:"column:station_sequence"`
	CumulativeTime  int    `gorm:"column:cumulative_time"`
}

type RouteStationListItem struct {
	StationID   string `gorm:"column:station_id;primaryKey"`
	RouteID     int    `gorm:"column:route_id;primaryKey"`
	StationName string `gorm:"column:station_name"`
}

func (RouteStation) TableName() string {
	return "subway_route_station"
}
