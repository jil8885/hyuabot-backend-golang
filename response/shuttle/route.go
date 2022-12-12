package shuttle

import (
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
)

type RouteListResponse struct {
	Route []RouteListItem `json:"route"`
}

type RouteListItem struct {
	Name        string           `json:"name"`
	Description RouteDescription `json:"description"`
}

type RouteDescription struct {
	Korean  string `json:"korean"`
	English string `json:"english"`
}

func CreateRouteListResponse(routeList []shuttle.RouteItem) RouteListResponse {
	var route []RouteListItem
	for _, routeItem := range routeList {
		route = append(route, RouteListItem{
			Name:        routeItem.Name,
			Description: RouteDescription{routeItem.DescriptionKorean, routeItem.DescriptionEnglish},
		})
	}
	return RouteListResponse{Route: route}
}
