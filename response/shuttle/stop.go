package shuttle

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
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
	Name          string      `json:"name"`
	Tag           string      `json:"tag"`
	RunningTime   RunningTime `json:"runningTime"`
	TimetableList []string    `json:"timetable"`
}

type StopRouteResponse struct {
	Name          string   `json:"name"`
	TimetableList []string `json:"timetable"`
}

type StopRouteTimetableResponse struct {
	Name     string   `json:"name"`
	Tag      string   `json:"tag"`
	Weekdays []string `json:"weekdays"`
	Weekends []string `json:"weekends"`
}

type StopRouteArrivalResponse struct {
	Name        string  `json:"name"`
	Tag         string  `json:"tag"`
	ArrivalList []int64 `json:"arrival"`
}

type RunningTime struct {
	Weekdays FirstLastTime `json:"weekdays"`
	Weekends FirstLastTime `json:"weekends"`
}

type FirstLastTime struct {
	FirstTime string `json:"first"`
	LastTime  string `json:"last"`
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

func CreateStopItemResponse(holidayType string, holiday bool, currentTime carbon.Carbon, stopItem shuttle.Stop) StopItemResponse {
	var routeStopList []StopRouteItem
	for _, routeStopItem := range stopItem.RouteList {
		routeStopList = append(routeStopList, CreateStopRouteItem(holidayType, holiday, currentTime, routeStopItem))
	}
	return StopItemResponse{
		Name:      stopItem.Name,
		Location:  StopLocation{stopItem.Latitude, stopItem.Longitude},
		RouteList: routeStopList,
	}
}

func CreateStopRouteItem(holidayType string, holiday bool, currentTime carbon.Carbon, routeStop shuttle.RouteStop) StopRouteItem {
	return StopRouteItem{
		Name:          routeStop.RouteName,
		Tag:           routeStop.ShuttleRoute.Tag,
		RunningTime:   CreateRunningTime(routeStop.TimetableList),
		TimetableList: CreateTimetable(holidayType, holiday, currentTime, routeStop.TimetableList),
	}
}

func CreateRunningTime(timetable []shuttle.Timetable) RunningTime {
	weekdaysFirstTime := ""
	weekdaysLastTime := ""
	weekendsFirstTime := ""
	weekendsLastTime := ""
	for _, timetableItem := range timetable {
		if timetableItem.Weekday {
			if weekdaysFirstTime == "" {
				weekdaysFirstTime = fmt.Sprintf(
					"%02d:%02d",
					timetableItem.DepartureTime.Microseconds/1000000/60/60,
					timetableItem.DepartureTime.Microseconds/1000000/60%60,
				)
			}
			weekdaysLastTime = fmt.Sprintf(
				"%02d:%02d",
				timetableItem.DepartureTime.Microseconds/1000000/60/60,
				timetableItem.DepartureTime.Microseconds/1000000/60%60,
			)
		} else {
			if weekendsFirstTime == "" {
				weekendsFirstTime = fmt.Sprintf(
					"%02d:%02d",
					timetableItem.DepartureTime.Microseconds/1000000/60/60,
					timetableItem.DepartureTime.Microseconds/1000000/60%60,
				)
			}
			weekendsLastTime = fmt.Sprintf(
				"%02d:%02d",
				timetableItem.DepartureTime.Microseconds/1000000/60/60,
				timetableItem.DepartureTime.Microseconds/1000000/60%60,
			)
		}
	}
	return RunningTime{
		Weekdays: FirstLastTime{FirstTime: weekdaysFirstTime, LastTime: weekdaysLastTime},
		Weekends: FirstLastTime{FirstTime: weekendsFirstTime, LastTime: weekendsLastTime},
	}
}

func CreateStopRouteArrivalItem(routeStop shuttle.RouteStop) StopRouteArrivalResponse {
	return StopRouteArrivalResponse{
		Name:        routeStop.RouteName,
		Tag:         routeStop.ShuttleRoute.Tag,
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
	date := carbon.Now().SetTime(0, 0, 0)
	return StopRouteTimetableResponse{
		Name:     routeStop.RouteName,
		Tag:      routeStop.ShuttleRoute.Tag,
		Weekdays: CreateTimetable("", true, date, weekdays),
		Weekends: CreateTimetable("", false, date, weekends),
	}
}
