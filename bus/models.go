package bus

type StopInfo struct {
	ID string `json:"id"`
	DepartureList []LineInfo `json:"departure_list"`
}

type LineInfo struct {
	Number string `json:"line_number"`
	DepartureList []DepartureItem `json:"departure_list"`
	TimeTableList []TimeTableItem `json:"timetable_list"`
}

type DepartureItem struct {
	Location int
	RemainedTime int
	RemainedSeat int
}

type TimeTableItem struct {
	DepartureTime string
}

type StopResponse struct {
	MsgBody stopMsgBody `xml:"msgBody"`
}

type stopMsgBody struct {
	BusArrivalList []busArrivalItem `xml:"busArrivalList"`
}

type RouteResponse struct {
	MsgBody routeMsgBody `xml:"msgBody"`
}

type routeMsgBody struct {
	BusArrivalItem busArrivalItem `xml:"busArrivalItem"`
}

type busArrivalItem struct {
	RouteID int `xml:"routeId"`
	LocationNo1    int `xml:"locationNo1"`
	LocationNo2    int `xml:"locationNo2"`
	PredictTime1   int `xml:"predictTime1"`
	PredictTime2   int `xml:"predictTime2"`
	RemainSeatCnt1 int `xml:"remainSeatCnt1"`
	RemainSeatCnt2 int `xml:"remainSeatCnt2"`
}

type BusTimeTableItem struct {
	Time string `json:"time"`
}

type BusTimeTableLine struct {
	Weekdays []BusTimeTableItem `json:"weekdays"`
	Sat []BusTimeTableItem      `json:"sat"`
	Sun []BusTimeTableItem      `json:"sun"`
}

type BusTimeTableJson struct {
	Line3102 BusTimeTableLine `json:"3102"`
	Line10_1 BusTimeTableLine `json:"10-1"`
	Line707_1 BusTimeTableLine `json:"707-1"`
}