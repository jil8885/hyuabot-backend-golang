package library

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/library"

type RoomListResponse struct {
	RoomList []RoomItemResponse `json:"room"`
}

type RoomItemResponse struct {
	RoomID    int    `json:"room_id"`
	Name      string `json:"name"`
	Total     int    `json:"total"`
	Active    int    `json:"active"`
	Occupied  int    `json:"occupied"`
	Available int    `json:"available"`
}

func CreateRoomListResponse(roomList []library.Room) RoomListResponse {
	var room []RoomItemResponse
	for _, roomItem := range roomList {
		room = append(room, CreateRoomItemResponse(roomItem))
	}
	return RoomListResponse{RoomList: room}
}

func CreateRoomItemResponse(roomItem library.Room) RoomItemResponse {
	return RoomItemResponse{
		RoomID:    roomItem.RoomID,
		Name:      roomItem.Name,
		Total:     roomItem.Total,
		Active:    roomItem.Active,
		Occupied:  roomItem.Occupied,
		Available: roomItem.Available,
	}
}
