package shuttle

type Route struct {
	Name               string      `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string      `gorm:"column:route_description_korean"`
	DescriptionEnglish string      `gorm:"column:route_description_english"`
	StopList           []RouteStop `gorm:"foreignKey:RouteName;references:Name"`
}

type RouteItem struct {
	Name               string `gorm:"column:route_name;primaryKey"`
	DescriptionKorean  string `gorm:"column:route_description_korean"`
	DescriptionEnglish string `gorm:"column:route_description_english"`
}

func (Route) TableName() string {
	return "shuttle_route"
}
