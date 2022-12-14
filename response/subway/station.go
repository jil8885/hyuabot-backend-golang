package subway

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/subway"

type StationListResponse struct {
	Station []StationListItem `json:"station"`
}

type StationListItem struct {
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
	RouteID     int    `json:"route_id"`
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
