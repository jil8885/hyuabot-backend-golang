package shuttle

type Route struct {
	Name               string      `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string      `gorm:"column:route_description_korean"`
	DescriptionEnglish string      `gorm:"column:route_description_english"`
	Tag                string      `gorm:"column:route_tag"`
	StartStopID        string      `gorm:"column:start_stop"`
	EndStopID          string      `gorm:"column:end_stop"`
	StartStop          Stop        `gorm:"foreignKey:StopID;references:StartStopID"`
	EndStop            Stop        `gorm:"foreignKey:StopID;references:EndStopID"`
	StopList           []RouteStop `gorm:"foreignKey:RouteName;references:Name"`
}

type RouteItem struct {
	Name               string `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string `gorm:"column:route_description_korean"`
	DescriptionEnglish string `gorm:"column:route_description_english"`
	Tag                string `gorm:"column:route_tag"`
	StartStopID        string `gorm:"column:start_stop"`
	EndStopID          string `gorm:"column:end_stop"`
	StartStop          Stop   `gorm:"foreignKey:StopID;references:StartStopID"`
	EndStop            Stop   `gorm:"foreignKey:StopID;references:EndStopID"`
}

func (Route) TableName() string {
	return "shuttle_route"
}
