package shuttle

import (
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
)

type StopListResponse struct {
	Stop []StopListItem `json:"stop"`
}

type StopListItem struct {
	Name     string       `json:"name"`
	Location StopLocation `json:"location"`
}

type StopItemResponse struct {
	Name      string          `json:"name"`
	Location  StopLocation    `json:"location"`
	RouteList []StopRouteItem `json:"route"`
}

type StopLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type StopRouteItem struct {
	Name          string   `json:"name"`
	TimetableList []string `json:"timetable"`
}

func CreateStopListResponse(stopList []shuttle.StopItem) StopListResponse {
	var stop []StopListItem
	for _, routeItem := range stopList {
		stop = append(stop, StopListItem{
			Name: routeItem.Name,
			Location: StopLocation{
				Latitude:  routeItem.Latitude,
				Longitude: routeItem.Longitude,
			},
		})
	}
	return StopListResponse{Stop: stop}
}

func CreateStopItemResponse(stopItem shuttle.Stop) StopItemResponse {
	var routeStopList []StopRouteItem
	for _, routeStopItem := range stopItem.RouteList {
		routeStopList = append(routeStopList, CreateStopRouteItem(routeStopItem.RouteName, routeStopItem.TimetableList))
	}
	return StopItemResponse{
		Name:      stopItem.Name,
		Location:  StopLocation{stopItem.Latitude, stopItem.Longitude},
		RouteList: routeStopList,
	}
}

func CreateStopRouteItem(RouteName string, timetableList []shuttle.Timetable) StopRouteItem {
	return StopRouteItem{
		Name:          RouteName,
		TimetableList: CreateTimetable(timetableList),
	}
}
