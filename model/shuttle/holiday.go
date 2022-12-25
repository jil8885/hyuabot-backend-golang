package shuttle

import "time"

type Holiday struct {
	Date         time.Time `gorm:"column:holiday_date;primaryKey"`
	HolidayType  string    `gorm:"column:holiday_type"`
	CalendarType string    `gorm:"column:calendar_type"`
}

func (Holiday) TableName() string {
	return "shuttle_holiday"
}
