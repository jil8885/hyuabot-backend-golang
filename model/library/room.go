package library

type Room struct {
	CampusID     int    `gorm:"column:campus_id;primaryKey"`
	RoomID       int    `gorm:"column:room_id;primaryKey"`
	Name         string `gorm:"column:room_name"`
	IsActive     bool   `gorm:"column:is_active"`
	IsReservable bool   `gorm:"column:is_reservable"`
	Total        int    `gorm:"column:total"`
	Active       int    `gorm:"column:active_total"`
	Occupied     int    `gorm:"column:occupied"`
	Available    int    `gorm:"column:available"`
}

func (Room) TableName() string {
	return "reading_room"
}
