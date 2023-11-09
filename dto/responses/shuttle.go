package responses

type ShuttleTimetableViewResponse struct {
	Data []ShuttleTimetableViewItem `json:"data"`
}

type ShuttleTimetableViewItem struct {
	Seq           int              `json:"seq"`
	PeriodType    string           `json:"period"`
	Weekday       bool             `json:"weekday"`
	Route         ShuttleRouteItem `json:"route"`
	StopName      string           `json:"stop"`
	DepartureTime string           `json:"time"`
}

type ShuttleRouteItem struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
