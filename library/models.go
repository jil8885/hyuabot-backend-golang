package library

type ReadingRoomInfo struct {
	Name string `json:"name"`
	IsActive     bool   `json:"isActive"`
	IsReservable bool   `json:"isReservable"`
	Total        int    `json:"total"`
	ActiveTotal  int    `json:"activeTotal"`
	Occupied     int    `json:"occupied"`
	Available    int    `json:"available"`
}

type ReadingRoomData struct {
	TotalCount int `json:"totalCount"`
	ReadingRoomList []ReadingRoomInfo `json:"list"`
}

type ReadingRoomJSON struct {
	Success bool `json:"success"`
	Data ReadingRoomData `json:"data"`
}