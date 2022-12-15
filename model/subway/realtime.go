package subway

import "time"

type Realtime struct {
	StationID         string       `gorm:"column:station_id;primaryKey"`
	ArrivalSequence   int          `gorm:"column:arrival_sequence;primaryKey"`
	Current           string       `gorm:"column:current_station_name"`
	RemainingTime     int          `gorm:"column:remaining_time"`
	RemainingStop     int          `gorm:"column:remaining_stop_count"`
	Heading           bool         `gorm:"column:up_down_type"`
	TerminalStationID string       `gorm:"column:terminal_station_id"`
	TerminalStation   RouteStation `gorm:"foreignKey:StationID;references:TerminalStationID"`
	TrainNumber       string       `gorm:"column:train_number"`
	LastUpdateTime    time.Time    `gorm:"column:last_updated_time"`
	IsExpress         bool         `gorm:"column:is_express_train"`
	Status            int          `gorm:"column:status_code"`
}

func (Realtime) TableName() string {
	return "subway_realtime"
}
