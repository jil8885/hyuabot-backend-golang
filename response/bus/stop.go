package bus

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"

type StopListResponse struct {
	Stop []StopListItem `json:"stop"`
}

type StopListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
		ID:   stop.StopID,
		Name: stop.StopName,
	}
}
