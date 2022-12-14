package bus

type Route struct {
	CompanyID        int    `gorm:"column:company_id"`
	CompanyName      string `gorm:"column:company_name"`
	CompanyTelephone string `gorm:"column:company_telephone"`
	DistrictCode     int    `gorm:"column:district_code"`
	UpFirstTime      string `gorm:"column:up_first_time"`
	UpLastTime       string `gorm:"column:up_last_time"`
	DownFirstTime    string `gorm:"column:down_first_time"`
	DownLastTime     string `gorm:"column:down_last_time"`
	StartStopID      int    `gorm:"column:start_stop_id"`
	StartStop        Stop   `gorm:"foreignKey:StopID;references:StartStopID"`
	EndStopID        int    `gorm:"column:end_stop_id"`
	EndStop          Stop   `gorm:"foreignKey:StopID;references:EndStopID"`
	RouteID          int    `gorm:"column:route_id;primaryKey"`
	RouteName        string `gorm:"column:route_name"`
	RouteTypeCode    int    `gorm:"column:route_type_code"`
	RouteTypeName    string `gorm:"column:route_type_name"`
}

func (Route) TableName() string {
	return "bus_route"
}
