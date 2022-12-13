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

//type RouteItemResponse struct {
//	Name        string           `json:"name"`
//	Description RouteDescription `json:"description"`
//	StopList    []RouteStopItem  `json:"stop_list"`
//}

type StopLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//type RouteStopItem struct {
//	Name          string   `json:"name"`
//	TimetableList []string `json:"timetable"`
//}
//
//type RouteLocationResponse struct {
//	Location []float64 `json:"location"`
//}

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

//func CreateRouteItemResponse(routeItem shuttle.Route) RouteItemResponse {
//	var routeStopList []RouteStopItem
//	for _, routeStopItem := range routeItem.StopList {
//		routeStopList = append(routeStopList, CreateRouteStopItem(routeStopItem.StopName, routeStopItem.TimetableList))
//	}
//	return RouteItemResponse{
//		Name:        routeItem.Name,
//		Description: RouteDescription{routeItem.DescriptionKorean, routeItem.DescriptionEnglish},
//		StopList:    routeStopList,
//	}
//}

//
//func CreateRouteStopItem(stopName string, timetableList []shuttle.Timetable) RouteStopItem {
//	return RouteStopItem{
//		Name:          stopName,
//		TimetableList: CreateTimetable(timetableList),
//	}
//}
//
//func CreateTimetable(timetableList []shuttle.Timetable) []string {
//	var timetable = make([]string, 0)
//	for _, timetableItem := range timetableList {
//		timetable = append(timetable, fmt.Sprintf(
//			"%02d:%02d",
//			timetableItem.DepartureTime.Microseconds/1000000/60/60,
//			timetableItem.DepartureTime.Microseconds/1000000/60%60,
//		))
//	}
//	sort.Slice(timetable, func(i, j int) bool {
//		return timetable[i] < timetable[j]
//	})
//	return timetable
//}
//
//func CreateRouteLocationResponse(routeItem shuttle.Route) RouteLocationResponse {
//	// 현재 시간 로딩 (KST)
//	loc, _ := time.LoadLocation("Asia/Seoul")
//	now := time.Now().In(loc)
//
//	cumulativeTimeList := make([]int, len(routeItem.StopList))
//	for i, routeStopItem := range routeItem.StopList {
//		cumulativeTimeList[i] = routeItem.StopList[len(routeItem.StopList)-1].CumulativeTime - routeStopItem.CumulativeTime
//	}
//	timetable := routeItem.StopList[len(routeItem.StopList)-1].TimetableList
//	remainingTime := 0
//	var location = make([]float64, 0)
//	for _, timetableItem := range timetable {
//		remainingTime = int((timetableItem.DepartureTime.Microseconds/1000000/60/60-int64(now.Hour()))*60 +
//			(timetableItem.DepartureTime.Microseconds/1000000/60%60 - int64(now.Minute())))
//		if remainingTime > cumulativeTimeList[0] {
//			continue
//		}
//		idx := len(cumulativeTimeList)
//		for i, cumulativeTime := range cumulativeTimeList {
//			if remainingTime > cumulativeTime || (remainingTime == cumulativeTime && i > 0) {
//				idx = i
//				break
//			}
//		}
//		location = append(location, float64(idx)-float64(remainingTime-cumulativeTimeList[idx])/float64(cumulativeTimeList[idx]-cumulativeTimeList[idx-1]))
//	}
//	sort.Slice(location, func(i, j int) bool {
//		return location[i] < location[j]
//	})
//	return RouteLocationResponse{
//		Location: location,
//	}
//}
