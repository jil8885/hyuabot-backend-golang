package shuttle

type RouteStop struct {
	RouteName      string      `gorm:"column:route_name;primaryKey"`
	ShuttleRoute   Route       `gorm:"foreignKey:RouteName;references:Name"`
	StopName       string      `gorm:"column:stop_name;primaryKey"`
	Order          int         `gorm:"column:stop_order"`
	CumulativeTime int         `gorm:"column:cumulative_time"`
	TimetableList  []Timetable `gorm:"foreignKey:RouteName,StopName;references:RouteName,StopName"`
}

func (RouteStop) TableName() string {
	return "shuttle_route_stop"
}
