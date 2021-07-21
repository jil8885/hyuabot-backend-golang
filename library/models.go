package library

type ReadingRoomInfo struct {
	Name string `json:"name" firestore:"name,omitempty"`
	IsActive     bool   `json:"isActive" firestore:"isActive,omitempty"`
	IsReservable bool   `json:"isReservable" firestore:"isReservable,omitempty"`
	Total        int    `json:"total" firestore:"total,omitempty"`
	ActiveTotal  int    `json:"activeTotal" firestore:"activeTotal,omitempty"`
	Occupied     int    `json:"occupied" firestore:"occupied,omitempty"`
	Available    int    `json:"available" firestore:"available,omitempty"`
}

type ReadingRoomData struct {
	TotalCount int `json:"totalCount"`
	ReadingRoomList []ReadingRoomInfo `json:"list"`
}

type ReadingRoomJSON struct {
	Success bool `json:"success"`
	Data ReadingRoomData `json:"data"`
}