package subway

import (
	"github.com/hyuabot-developers/hyuabot-backend-golang/model/subway"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StationListResponse struct {
	Station []StationListItem `json:"station"`
}

type StationListItem struct {
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
	RouteID     int    `json:"route_id"`
}

type StationItemResponse struct {
	StationID   string                       `json:"stationID"`
	StationName string                       `json:"stationName"`
	RouteID     int                          `json:"routeID"`
	Realtime    StationRealtimeHeadingGroup  `json:"realtime"`
	Timetable   StationTimetableHeadingGroup `json:"timetable"`
}

type StationRealtimeHeadingGroup struct {
	Up   []StationRealtimeItem `json:"up"`
	Down []StationRealtimeItem `json:"down"`
}

type StationRealtimeItem struct {
	TerminalStationName string `json:"terminal"`
	TrainNo             string `json:"trainNumber"`
	RemainingTime       int    `json:"time"`
	RemainingStop       int    `json:"stop"`
	IsExpress           bool   `json:"express"`
	IsUp                bool   `json:"heading"`
	CurrentLocation     string `json:"location"`
}

type StationTimetableHeadingGroup struct {
	Up   []StationTimetableItem `json:"up"`
	Down []StationTimetableItem `json:"down"`
}

type StationTimetableItem struct {
	TerminalStationName string `json:"terminal"`
	DepartureTime       string `json:"departureTime"`
}

func CreateStationListResponse(stationList []subway.RouteStationListItem) StationListResponse {
	station := make([]StationListItem, 0)
	for _, stationItem := range stationList {
		station = append(station, CreateStationListItem(stationItem))
	}
	return StationListResponse{Station: station}
}

func CreateStationListItem(stationItem subway.RouteStationListItem) StationListItem {
	return StationListItem{
		StationID:   stationItem.StationID,
		StationName: stationItem.StationName,
		RouteID:     stationItem.RouteID,
	}
}

func CreateStationItemResponse(stationItem subway.RouteStationItem) StationItemResponse {
	realtime, maxUP, maxDown := CreateStationRealtimeGroup(stationItem.RealtimeList)
	return StationItemResponse{
		StationID:   stationItem.StationID,
		StationName: stationItem.StationName,
		RouteID:     stationItem.RouteID,
		Realtime:    realtime,
		Timetable:   CreateStationTimetableGroup(stationItem.TimetableList, maxUP, maxDown),
	}
}

func CreateStationRealtimeGroup(realtimeList []subway.Realtime) (StationRealtimeHeadingGroup, int, int) {
	sort.Slice(realtimeList, func(i, j int) bool {
		return realtimeList[i].ArrivalSequence < realtimeList[j].ArrivalSequence
	})
	var up = make([]StationRealtimeItem, 0)
	var down = make([]StationRealtimeItem, 0)
	var maxUp = 0
	var maxDown = 0
	for _, realtimeItem := range realtimeList {
		if realtimeItem.Heading {
			up = append(up, CreateStationRealtimeItem(realtimeItem))
			if maxUp < realtimeItem.RemainingTime {
				maxUp = realtimeItem.RemainingTime
			}
		} else {
			down = append(down, CreateStationRealtimeItem(realtimeItem))
			if maxDown < realtimeItem.RemainingTime {
				maxDown = realtimeItem.RemainingTime
			}
		}
	}
	return StationRealtimeHeadingGroup{Up: up, Down: down}, maxUp, maxDown
}

func CreateStationRealtimeItem(realtimeItem subway.Realtime) StationRealtimeItem {
	return StationRealtimeItem{
		TerminalStationName: realtimeItem.TerminalStation.StationName,
		TrainNo:             realtimeItem.TrainNumber,
		RemainingTime:       realtimeItem.RemainingTime,
		RemainingStop:       realtimeItem.RemainingStop,
		IsExpress:           realtimeItem.IsExpress,
		IsUp:                realtimeItem.Heading,
		CurrentLocation:     realtimeItem.Current,
	}
}

func CreateStationTimetableGroup(timetableList []subway.Timetable, maxUp int, maxDown int) StationTimetableHeadingGroup {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	sort.Slice(timetableList, func(i, j int) bool {
		return timetableList[i].DepartureTime < timetableList[j].DepartureTime
	})
	var up = make([]StationTimetableItem, 0)
	var down = make([]StationTimetableItem, 0)
	for _, timetableItem := range timetableList {
		hour, _ := strconv.Atoi(strings.Split(timetableItem.DepartureTime, ":")[0])
		minute, _ := strconv.Atoi(strings.Split(timetableItem.DepartureTime, ":")[1])
		remainingTime := (hour-now.Hour())*60 + (minute - now.Minute())
		if timetableItem.Heading == "up" && remainingTime > maxUp {
			up = append(up, CreateStationTimetableItem(timetableItem))
		} else if timetableItem.Heading == "down" && remainingTime > maxDown {
			down = append(down, CreateStationTimetableItem(timetableItem))
		}
	}
	return StationTimetableHeadingGroup{Up: up, Down: down}
}

func CreateStationTimetableItem(timetableItem subway.Timetable) StationTimetableItem {
	return StationTimetableItem{
		TerminalStationName: timetableItem.TerminalStation.StationName,
		DepartureTime:       timetableItem.DepartureTime,
	}
}
