package shuttle

import (
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
)

type StopTimetableListResponse struct {
	Stop []StopRouteTimetableItem `json:"stop"`
}

type StopRouteTimetableItem struct {
	Name string `json:"name"`
	Route []StopRouteTimetableResponse `json:"route"`
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
			Name: routeItem.RouteName,
			Weekdays: CreateTimetable(weekdays),
			Weekends: CreateTimetable(weekends),
		})
	}
	return StopRouteTimetableItem{
		Name: stop.Name,
		Route: route,
	}
}
