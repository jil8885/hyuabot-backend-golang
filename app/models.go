package app

import (
	"github.com/jil8885/hyuabot-backend-golang/bus"
	"github.com/jil8885/hyuabot-backend-golang/library"
	"github.com/jil8885/hyuabot-backend-golang/shuttle"
	"github.com/jil8885/hyuabot-backend-golang/subway"
)

type ShuttleStopRequest struct {
	BusStop string `json:"busStop"`
}

type CampusRequest struct {
	Campus string `json:"campus"`
}

type BusRouteRequest struct {
	Route string `json:"routeID"`
}

type ShuttleDepartureByStop struct {
	BusForStation  []shuttle.Departure `json:"busForStation"`
	BusForTerminal []shuttle.Departure `json:"busForTerminal"`
}

type ShuttleStop struct {
	RoadViewLink string `json:"roadViewLink"`
	FirstBusForStation string `json:"firstBusForStation"`
	LastBusForStation string `json:"lastBusForStation"`
	FirstBusForTerminal string `json:"firstBusForTerminal"`
	LastBusForTerminal string `json:"lastBusForTerminal"`
	Weekdays ShuttleDepartureByStop `json:"weekdays"`
	Weekends ShuttleDepartureByStop `json:"weekends"`
}

type SubwayDepartureSeoul struct {
	Line2 subway.RealtimeDataResult `json:"main"`
}

type SubwayDepartureERICA struct {
	Line4 subway.RealtimeDataResult `json:"main"`
	LineSuin subway.TimetableDataResult `json:"sub"`
}

type Bus struct {
	Realtime map[string][]bus.DepartureItem `json:"realtime"`
	Timetable bus.BusTimeTableJson `json:"timetable"`
}

type BusByRoute struct {
	Realtime []bus.DepartureItem `json:"realtime"`
	Timetable bus.BusTimeTableLine `json:"timetable"`
}

type ReadingRoomByCampus struct {
	OpenedRooms []library.ReadingRoomInfo `json:"rooms"`
}