package shuttle

import "time"

type Period struct {
	Type  string    `gorm:"column:period_type;primaryKey"`
	Start time.Time `gorm:"column:period_start;primaryKey"`
	End   time.Time `gorm:"column:period_end;primaryKey"`
}

func (Period) TableName() string {
	return "shuttle_period"
}
