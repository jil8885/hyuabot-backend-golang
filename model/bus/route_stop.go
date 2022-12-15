package bus

type RouteStop struct {
	RouteID       int         `gorm:"column:route_id;primaryKey"`
	RouteItem     Route       `gorm:"foreignKey:RouteID;references:RouteID"`
	StopID        int         `gorm:"column:stop_id;primaryKey"`
	StopSequence  int         `gorm:"column:stop_sequence"`
	StartStopID   int         `gorm:"column:start_stop_id"`
	StartStop     Stop        `gorm:"foreignKey:StartStopID;references:StopID"`
	TimetableList []Timetable `gorm:"foreignKey:RouteID,StartStopID;references:RouteID,StartStopID"`
	RealtimeList  []Realtime  `gorm:"foreignKey:RouteID,StopID;references:RouteID,StopID"`
}

func (RouteStop) TableName() string {
	return "bus_route_stop"
}
