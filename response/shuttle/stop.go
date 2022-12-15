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

type StopRouteResponse struct {
	Name string `json:"name"`
	TimetableList []string `json:"timetable"`
}

type StopRouteTimetableResponse struct {
	Name string `json:"name"`
	Weekdays []string `json:"weekdays"`
	Weekends []string `json:"weekends"`
}

type StopRouteArrivalResponse struct {
	Name string `json:"name"`
	ArrivalList []int64 `json:"arrival"`
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
		routeStopList = append(routeStopList, CreateStopRouteItem(routeStopItem))
	}
	return StopItemResponse{
		Name:      stopItem.Name,
		Location:  StopLocation{stopItem.Latitude, stopItem.Longitude},
		RouteList: routeStopList,
	}
}

func CreateStopRouteItem(routeStop shuttle.RouteStop) StopRouteItem {
	return StopRouteItem{
		Name:          routeStop.RouteName,
		TimetableList: CreateTimetable(routeStop.TimetableList),
	}
}

func CreateStopRouteArrivalItem(routeStop shuttle.RouteStop) StopRouteArrivalResponse {
	return StopRouteArrivalResponse{
		Name:          routeStop.RouteName,
		ArrivalList: CreateArrival(routeStop.TimetableList),
	}
}

func CreateStopRouteTimetableResponse(routeStop shuttle.RouteStop) StopRouteTimetableResponse {
	var weekdays = make([]shuttle.Timetable, 0)
	var weekends = make([]shuttle.Timetable, 0)
	for _, timetable := range routeStop.TimetableList {
		if timetable.Weekday {
			weekdays = append(weekdays, timetable)
		} else {
			weekends = append(weekends, timetable)
		}
	}
	return StopRouteTimetableResponse{
		Name:          routeStop.RouteName,
		Weekdays: CreateTimetable(weekdays),
		Weekends: CreateTimetable(weekends),
	}
}

