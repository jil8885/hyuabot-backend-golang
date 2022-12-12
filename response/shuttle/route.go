package shuttle

import (
	"fmt"
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	"sort"
)

type RouteListResponse struct {
	Route []RouteListItem `json:"route"`
}

type RouteListItem struct {
	Name        string           `json:"name"`
	Description RouteDescription `json:"description"`
}

type RouteItemResponse struct {
	Name        string           `json:"name"`
	Description RouteDescription `json:"description"`
	StopList    []RouteStopItem  `json:"stop_list"`
}

type RouteDescription struct {
	Korean  string `json:"korean"`
	English string `json:"english"`
}

type RouteStopItem struct {
	Name          string   `json:"name"`
	TimetableList []string `json:"timetable"`
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

func CreateRouteItemResponse(routeItem shuttle.Route) RouteItemResponse {
	var routeStopList []RouteStopItem
	for _, routeStopItem := range routeItem.StopList {
		routeStopList = append(routeStopList, CreateRouteStopItem(routeStopItem.StopName, routeStopItem.TimetableList))
	}
	return RouteItemResponse{
		Name:        routeItem.Name,
		Description: RouteDescription{routeItem.DescriptionKorean, routeItem.DescriptionEnglish},
		StopList:    routeStopList,
	}
}

func CreateRouteStopItem(stopName string, timetableList []shuttle.Timetable) RouteStopItem {
	return RouteStopItem{
		Name:          stopName,
		TimetableList: CreateTimetable(timetableList),
	}
}

func CreateTimetable(timetableList []shuttle.Timetable) []string {
	var timetable []string
	for _, timetableItem := range timetableList {
		timetable = append(timetable, fmt.Sprintf(
			"%02d:%02d",
			timetableItem.DepartureTime.Microseconds/1000000/60/60,
			timetableItem.DepartureTime.Microseconds/1000000/60%60,
		))
	}
	sort.Slice(timetable, func(i, j int) bool {
		return timetable[i] < timetable[j]
	})
	return timetable
}
