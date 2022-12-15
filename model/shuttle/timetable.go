package shuttle

import (
	"github.com/jackc/pgtype"
)

type Timetable struct {
	PeriodType    string      `gorm:"column:period_type;primaryKey"`
	Weekday       bool        `gorm:"column:weekday;primaryKey"`
	RouteName     string      `gorm:"column:route_name;primaryKey"`
	StopName      string      `gorm:"column:stop_name;primaryKey"`
	DepartureTime pgtype.Time `gorm:"column:departure_time;primaryKey"`
}

func (Timetable) TableName() string {
	return "shuttle_timetable"
}
