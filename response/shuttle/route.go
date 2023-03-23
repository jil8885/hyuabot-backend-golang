package shuttle

import (
	"fmt"
	"sort"
	"time"

	"github.com/golang-module/carbon/v2"

	"github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
)

type RouteListResponse struct {
	Route []RouteListItem `json:"route"`
}

type RouteListItem struct {
	Name        string           `json:"name"`
	Tag         string           `json:"tag"`
	Description RouteDescription `json:"description"`
}

type RouteItemResponse struct {
	Name        string           `json:"name"`
	Tag         string           `json:"tag"`
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

type RouteLocationResponse struct {
	Location []float64 `json:"location"`
}

func CreateRouteListResponse(routeList []shuttle.RouteItem) RouteListResponse {
	var route []RouteListItem
	for _, routeItem := range routeList {
		route = append(route, RouteListItem{
			Name:        routeItem.Name,
			Tag:         routeItem.Tag,
			Description: RouteDescription{routeItem.DescriptionKorean, routeItem.DescriptionEnglish},
		})
	}
	return RouteListResponse{Route: route}
}

func CreateRouteItemResponse(holidayType string, holiday bool, currentTime carbon.Carbon, routeItem shuttle.Route) RouteItemResponse {
	var routeStopList []RouteStopItem
	for _, routeStopItem := range routeItem.StopList {
		routeStopList = append(routeStopList, CreateRouteStopItem(holidayType, holiday, currentTime, routeStopItem.StopName, routeStopItem.TimetableList))
	}
	return RouteItemResponse{
		Name:        routeItem.Name,
		Tag:         routeItem.Tag,
		Description: RouteDescription{routeItem.DescriptionKorean, routeItem.DescriptionEnglish},
		StopList:    routeStopList,
	}
}

func CreateRouteStopItem(holidayType string, holiday bool, currentTime carbon.Carbon, stopName string, timetableList []shuttle.Timetable) RouteStopItem {
	return RouteStopItem{
		Name:          stopName,
		TimetableList: CreateTimetable(holidayType, holiday, currentTime, timetableList),
	}
}

func CreateTimetable(holidayType string, holiday bool, currentTime carbon.Carbon, timetableList []shuttle.Timetable) []string {
	var timetable = make([]string, 0)
	if holidayType != "halt" {
		for _, timetableItem := range timetableList {
			if timetableItem.Weekday != holiday ||
				fmt.Sprintf("%02d:%02d", timetableItem.DepartureTime.Microseconds/1000000/60/60, timetableItem.DepartureTime.Microseconds/1000000/60%60) < currentTime.ToTimeString()[0:5] {
				continue
			}
			timetable = append(timetable, fmt.Sprintf(
				"%02d:%02d",
				timetableItem.DepartureTime.Microseconds/1000000/60/60,
				timetableItem.DepartureTime.Microseconds/1000000/60%60,
			))
		}
	}
	sort.Slice(timetable, func(i, j int) bool {
		return timetable[i] < timetable[j]
	})
	return timetable
}

func CreateArrival(timetableList []shuttle.Timetable) []int64 {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var arrival = make([]int64, 0)
	for _, timetableItem := range timetableList {
		arrival = append(arrival,
			(timetableItem.DepartureTime.Microseconds/1000000/60/60-int64(now.Hour()))*60+
				timetableItem.DepartureTime.Microseconds/1000000/60%60-int64(now.Minute()),
		)
	}
	sort.Slice(arrival, func(i, j int) bool {
		return arrival[i] < arrival[j]
	})
	return arrival
}

func CreateRouteLocationResponse(holidayType string, routeItem shuttle.Route) RouteLocationResponse {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	cumulativeTimeList := make([]int, len(routeItem.StopList))
	for i, routeStopItem := range routeItem.StopList {
		cumulativeTimeList[i] = routeItem.StopList[len(routeItem.StopList)-1].CumulativeTime - routeStopItem.CumulativeTime
	}
	timetable := routeItem.StopList[len(routeItem.StopList)-1].TimetableList
	var location = make([]float64, 0)
	if holidayType != "halt" {
		for _, timetableItem := range timetable {
			remainingTime := int((timetableItem.DepartureTime.Microseconds/1000000/60/60-int64(now.Hour()))*60 +
				(timetableItem.DepartureTime.Microseconds/1000000/60%60 - int64(now.Minute())))
			if remainingTime > cumulativeTimeList[0] {
				continue
			}
			idx := len(cumulativeTimeList)
			for i, cumulativeTime := range cumulativeTimeList {
				if remainingTime > cumulativeTime || (remainingTime == cumulativeTime && i > 0) {
					idx = i
					break
				}
			}
			location = append(location, float64(idx)-float64(remainingTime-cumulativeTimeList[idx])/float64(cumulativeTimeList[idx]-cumulativeTimeList[idx-1]))
		}
	}
	sort.Slice(location, func(i, j int) bool {
		return location[i] < location[j]
	})
	return RouteLocationResponse{
		Location: location,
	}
}
