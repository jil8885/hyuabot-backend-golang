package bus

import "time"

type Realtime struct {
	StopID             int       `gorm:"column:stop_id;primaryKey"`
	RouteID            int       `gorm:"column:route_id;primaryKey"`
	ArrivalSequence    int       `gorm:"column:arrival_sequence"`
	RemainingStopCount int       `gorm:"column:remaining_stop_count"`
	RemainingSeatCount int       `gorm:"column:remaining_seat_count"`
	RemainingTime      int       `gorm:"column:remaining_time"`
	LowPlate           bool      `gorm:"column:low_plate"`
	LastUpdateTime     time.Time `gorm:"column:last_updated_time"`
}

func (Realtime) TableName() string {
	return "bus_realtime"
}
