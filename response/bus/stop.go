package bus

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"

type StopListResponse struct {
	Stop []StopListItem `json:"stop"`
}

type StopListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StopItemResponse struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	MobileNumber string          `json:"mobileNumber"`
	Location     StopLocation    `json:"location"`
	Route        []StopRouteItem `json:"route"`
}

type StopLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type StopRouteItem struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Start        StartStop      `json:"start"`
	RealtimeList []RealtimeItem `json:"realtime"`
}

type StartStop struct {
	StopID        int      `json:"id"`
	StopName      string   `json:"name"`
	TimetableList []string `json:"timetable"`
}

type RealtimeItem struct {
	RemainingStop int  `json:"stop"`
	RemainingTime int  `json:"time"`
	RemainingSeat int  `json:"seat"`
	LowPlate      bool `json:"lowPlate"`
}

type RouteTimetableResponse struct {
	Weekdays []string `json:"weekdays"`
	Saturday []string `json:"saturday"`
	Sunday   []string `json:"sunday"`
}

func CreateStopListResponse(stopList []bus.Stop) StopListResponse {
	var stopListItems []StopListItem
	for _, stop := range stopList {
		stopListItems = append(stopListItems, CreateStopListItem(stop))
	}
	return StopListResponse{Stop: stopListItems}
}

func CreateStopListItem(stop bus.Stop) StopListItem {
	return StopListItem{
		Name: stop.StopName,
		ID:   stop.StopID,
	}
}

func CreateStopItemResponse(stop bus.Stop) StopItemResponse {
	var stopRouteItems []StopRouteItem
	for _, routeStop := range stop.RouteList {
		stopRouteItems = append(stopRouteItems, CreateStopRouteItem(routeStop))
	}
	return StopItemResponse{
		ID:           stop.StopID,
		Name:         stop.StopName,
		MobileNumber: stop.MobileNumber,
		Location: StopLocation{
			Latitude:  stop.Latitude,
			Longitude: stop.Longitude,
		},
		Route: stopRouteItems,
	}
}

func CreateStopRouteItem(routeStop bus.RouteStop) StopRouteItem {
	var realtimeList = make([]RealtimeItem, 0)
	for _, realtime := range routeStop.RealtimeList {
		realtimeList = append(realtimeList, CreateRealtimeItem(realtime))
	}
	return StopRouteItem{
		ID:           routeStop.RouteID,
		Name:         routeStop.RouteItem.RouteName,
		Start:        CreateStartStop(routeStop.StartStop, routeStop.TimetableList),
		RealtimeList: realtimeList,
	}
}

func CreateStartStop(stop bus.Stop, timetable []bus.Timetable) StartStop {
	var timetableList = make([]string, 0)
	for _, time := range timetable {
		timetableList = append(timetableList, time.DepartureTime)
	}
	return StartStop{
		StopID:        stop.StopID,
		StopName:      stop.StopName,
		TimetableList: timetableList,
	}
}

func CreateRealtimeItem(realtime bus.Realtime) RealtimeItem {
	return RealtimeItem{
		RemainingStop: realtime.RemainingStopCount,
		RemainingTime: realtime.RemainingTime,
		RemainingSeat: realtime.RemainingSeatCount,
		LowPlate:      realtime.LowPlate,
	}
}

func CreateRouteTimetableResponse(timetable []bus.Timetable) RouteTimetableResponse {
	var weekdays = make([]string, 0)
	var saturday = make([]string, 0)
	var sunday = make([]string, 0)
	for _, time := range timetable {
		switch time.Weekday {
		case "weekdays":
			weekdays = append(weekdays, time.DepartureTime)
		case "saturday":
			saturday = append(saturday, time.DepartureTime)
		case "sunday":
			sunday = append(sunday, time.DepartureTime)
		}
	}
	return RouteTimetableResponse{
		Weekdays: weekdays,
		Saturday: saturday,
		Sunday:   sunday,
	}
}
