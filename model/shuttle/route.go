package shuttle

type Route struct {
	Name               string      `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string      `gorm:"column:route_description_korean"`
	DescriptionEnglish string      `gorm:"column:route_description_english"`
	Tag                string      `gorm:"column:route_tag"`
	StartStopName      string      `gorm:"column:start_stop"`
	EndStopName        string      `gorm:"column:end_stop"`
	StartStop          Stop        `gorm:"foreignKey:StartStopName;references:Name"`
	EndStop            Stop        `gorm:"foreignKey:EndStopName;references:Name"`
	StopList           []RouteStop `gorm:"foreignKey:RouteName;references:Name"`
}

type RouteItem struct {
	Name               string `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string `gorm:"column:route_description_korean"`
	DescriptionEnglish string `gorm:"column:route_description_english"`
	Tag                string `gorm:"column:route_tag"`
	StartStopName      string `gorm:"column:start_stop"`
	EndStopName        string `gorm:"column:end_stop"`
	StartStop          Stop   `gorm:"foreignKey:StartStopName;references:Name"`
	EndStop            Stop   `gorm:"foreignKey:EndStopName;references:Name"`
}

func (Route) TableName() string {
	return "shuttle_route"
}
