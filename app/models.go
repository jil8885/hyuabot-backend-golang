package app

import (
	"github.com/jil8885/hyuabot-backend-golang/bus"
	"github.com/jil8885/hyuabot-backend-golang/library"
	"github.com/jil8885/hyuabot-backend-golang/shuttle"
	"github.com/jil8885/hyuabot-backend-golang/subway"
)

type ShuttleStopRequest struct {
	BusStop string `query:"stop" json:"busStop"`
}

type CampusRequest struct {
	Campus string `query:"campus" json:"campus"`
}

type BusRouteRequest struct {
	Route string `query:"route_id" json:"routeID"`
}

type ShuttleDepartureByStop struct {
	BusForStation  []shuttle.Departure `json:"busForStation"`
	BusForTerminal []shuttle.Departure `json:"busForTerminal"`
}

type ShuttleStop struct {
	RoadViewLink        string                 `json:"roadViewLink"`
	FirstBusForStation  string                 `json:"firstBusForStation"`
	LastBusForStation   string                 `json:"lastBusForStation"`
	FirstBusForTerminal string                 `json:"firstBusForTerminal"`
	LastBusForTerminal  string                 `json:"lastBusForTerminal"`
	Weekdays            ShuttleDepartureByStop `json:"weekdays"`
	Weekends            ShuttleDepartureByStop `json:"weekends"`
}

type SubwayDepartureSeoul struct {
	Line2 subway.RealtimeDataResult `json:"main"`
}

type SubwayDepartureByLine struct {
	RealtimeList  subway.RealtimeDataResult  `json:"realtime"`
	TimetableList subway.TimetableDataResult `json:"timetable"`
}

type SubwayDepartureERICA struct {
	Line4    SubwayDepartureByLine `json:"main"`
	LineSuin SubwayDepartureByLine `json:"sub"`
}

type Bus struct {
	LineRed            BusByRoute `json:"3102"`
	LineBlue           BusByRoute `json:"707-1"`
	LineGreenToStation BusByRoute `json:"10-1_station"`
	LineGreenToCampus  BusByRoute `json:"10-1_campus"`
}

type BusByRoute struct {
	Realtime  []bus.DepartureItem  `json:"realtime"`
	Timetable bus.BusTimeTableLine `json:"timetable"`
}

type ReadingRoomByCampus struct {
	OpenedRooms []library.ReadingRoomInfo `json:"rooms"`
}
