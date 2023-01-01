package shuttle

import (
	"github.com/golang-module/carbon/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
)

type StopTimetableListResponse struct {
	Stop []StopRouteTimetableItem `json:"stop"`
}

type StopRouteTimetableItem struct {
	Name  string                       `json:"name"`
	Route []StopRouteTimetableResponse `json:"route"`
}

type StopArrivalListResponse struct {
	Stop []StopRouteArrivalItem `json:"stop"`
}

type StopRouteArrivalItem struct {
	Name  string                     `json:"name"`
	Route []StopRouteArrivalResponse `json:"route"`
}

func CreateStopTimetableListResponse(stopList []model.Stop) StopTimetableListResponse {
	var stop = make([]StopRouteTimetableItem, 0)
	for _, stopItem := range stopList {
		stop = append(stop, CreateStopTimetableItem(stopItem))
	}
	return StopTimetableListResponse{Stop: stop}
}

func CreateStopTimetableItem(stop model.Stop) StopRouteTimetableItem {
	var route = make([]StopRouteTimetableResponse, 0)
	date := carbon.Now().SetTime(0, 0, 0)
	for _, routeItem := range stop.RouteList {
		var weekdays = make([]model.Timetable, 0)
		var weekends = make([]model.Timetable, 0)
		for _, timetableItem := range routeItem.TimetableList {
			if timetableItem.Weekday {
				weekdays = append(weekdays, timetableItem)
			} else {
				weekends = append(weekends, timetableItem)
			}
		}
		route = append(route, StopRouteTimetableResponse{
			Name:     routeItem.RouteName,
			Tag:      routeItem.ShuttleRoute.Tag,
			Weekdays: CreateTimetable("", true, date, weekdays),
			Weekends: CreateTimetable("", false, date, weekends),
		})
	}
	return StopRouteTimetableItem{
		Name:  stop.Name,
		Route: route,
	}
}

func CreateStopArrivalListResponse(holidayType string, stopList []model.Stop) StopArrivalListResponse {
	var stop = make([]StopRouteArrivalItem, 0)
	for _, stopItem := range stopList {
		stop = append(stop, CreateStopArrivalItem(holidayType, stopItem))
	}
	return StopArrivalListResponse{Stop: stop}
}

func CreateStopArrivalItem(holidayType string, stop model.Stop) StopRouteArrivalItem {
	var route = make([]StopRouteArrivalResponse, 0)
	for _, routeItem := range stop.RouteList {
		if holidayType == "halt" {
			route = append(route, StopRouteArrivalResponse{
				Name:        routeItem.RouteName,
				Tag:         routeItem.ShuttleRoute.Tag,
				ArrivalList: make([]int64, 0),
			})
		} else {
			route = append(route, StopRouteArrivalResponse{
				Name:        routeItem.RouteName,
				Tag:         routeItem.ShuttleRoute.Tag,
				ArrivalList: CreateArrival(routeItem.TimetableList),
			})
		}
	}
	return StopRouteArrivalItem{
		Name:  stop.Name,
		Route: route,
	}
}
