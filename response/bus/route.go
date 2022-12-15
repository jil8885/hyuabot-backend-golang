package bus

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"

type RouteListResponse struct {
	Route []RouteListItem `json:"route"`
}

type RouteListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RouteItemResponse struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Type        RouteType             `json:"type"`
	Company     RouteCompany          `json:"company"`
	RunningTime RouteRunningTimeGroup `json:"runningTime"`
	Start       RouteStop             `json:"start"`
	End         RouteStop             `json:"end"`
}

type RouteCompany struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type RouteRunningTimeGroup struct {
	Up   RouteRunningTime `json:"up"`
	Down RouteRunningTime `json:"down"`
}

type RouteRunningTime struct {
	FirstTime string `json:"first"`
	LastTime  string `json:"last"`
}

type RouteStop struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RouteType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateRouteListResponse(routeList []bus.Route) RouteListResponse {
	var routeListItems []RouteListItem
	for _, route := range routeList {
		routeListItems = append(routeListItems, CreateRouteListItem(route))
	}
	return RouteListResponse{Route: routeListItems}
}

func CreateRouteListItem(route bus.Route) RouteListItem {
	return RouteListItem{
		ID:   route.RouteID,
		Name: route.RouteName,
	}
}

func CreateRouteType(route bus.Route) RouteType {
	return RouteType{
		ID:   route.RouteTypeCode,
		Name: route.RouteTypeName,
	}
}

func CreateRouteCompany(route bus.Route) RouteCompany {
	return RouteCompany{
		ID:    route.CompanyID,
		Name:  route.CompanyName,
		Phone: route.CompanyTelephone,
	}
}

func CreateRouteRunningTimeGroup(route bus.Route) RouteRunningTimeGroup {
	return RouteRunningTimeGroup{
		Up:   CreateRouteRunningTime(route.UpFirstTime, route.UpLastTime),
		Down: CreateRouteRunningTime(route.DownFirstTime, route.DownLastTime),
	}
}

func CreateRouteRunningTime(firstTime string, lastTime string) RouteRunningTime {
	return RouteRunningTime{
		FirstTime: firstTime,
		LastTime:  lastTime,
	}
}

func CreateRouteStop(stop bus.Stop) RouteStop {
	return RouteStop{
		ID:   stop.StopID,
		Name: stop.StopName,
	}
}

func CreateRouteItemResponse(route bus.Route) RouteItemResponse {
	return RouteItemResponse{
		ID:          route.RouteID,
		Name:        route.RouteName,
		Type:        CreateRouteType(route),
		Company:     CreateRouteCompany(route),
		RunningTime: CreateRouteRunningTimeGroup(route),
		Start:       CreateRouteStop(route.StartStop),
		End:         CreateRouteStop(route.EndStop),
	}
}
