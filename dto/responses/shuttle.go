package responses

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type ShuttleTimetableViewResponse struct {
	Data []ShuttleTimetableViewItem `json:"data"`
}

type ShuttleTimetableViewItem struct {
	Seq           int                       `json:"seq"`
	PeriodType    string                    `json:"period"`
	Weekday       bool                      `json:"weekday"`
	Route         ShuttleTimetableRouteItem `json:"route"`
	StopName      string                    `json:"stop"`
	DepartureTime string                    `json:"time"`
}

type ShuttleTimetableRouteItem struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type ShuttleTimetableResponse struct {
	Data []ShuttleTimetableItem `json:"data"`
}

type ShuttleTimetableItem struct {
	Seq           int    `json:"seq"`
	PeriodType    string `json:"period"`
	Weekday       bool   `json:"weekday"`
	Route         string `json:"route"`
	DepartureTime string `json:"time"`
}

type ShuttleRouteListResponse struct {
	Data []ShuttleRouteItem `json:"data"`
}

type ShuttleRouteItem struct {
	Name        string                  `json:"name"`
	Tag         string                  `json:"tag"`
	Description ShuttleRouteDescription `json:"description"`
	Start       ShuttleStopItem         `json:"start"`
	End         ShuttleStopItem         `json:"end"`
}

type ShuttleRouteDescription struct {
	Korean  string `json:"korean"`
	English string `json:"english"`
}

type ShuttleStopItem struct {
	Name      string       `json:"name"`
	Latitude  null.Float64 `json:"latitude"`
	Longitude null.Float64 `json:"longitude"`
}

type ShuttleRouteDetailItem struct {
	Name        string                  `json:"name"`
	Tag         string                  `json:"tag"`
	Description ShuttleRouteDescription `json:"description"`
	Start       ShuttleStopItem         `json:"start"`
	End         ShuttleStopItem         `json:"end"`
	Stops       []ShuttleRouteStopItem  `json:"stops"`
}

type ShuttleStopListResponse struct {
	Data []ShuttleStopItem `json:"data"`
}

type ShuttleRouteStopListResponse struct {
	Data []ShuttleRouteStopItem `json:"data"`
}

type ShuttleRouteStopItem struct {
	Name           string        `json:"name"`
	Seq            int           `json:"seq"`
	CumulativeTime time.Duration `json:"cumulativeTime"`
}

type ShuttleRouteStopDetailItem struct {
	Route          string        `json:"route"`
	Stop           string        `json:"stop"`
	Seq            int           `json:"seq"`
	CumulativeTime time.Duration `json:"cumulativeTime"`
}

type ShuttlePeriodListResponse struct {
	Data []ShuttlePeriodItem `json:"data"`
}

type ShuttlePeriodItem struct {
	PeriodType string `json:"type"`
	StartDate  string `json:"start"`
	EndDate    string `json:"end"`
}

type ShuttleHolidayListResponse struct {
	Data []ShuttleHolidayItem `json:"data"`
}

type ShuttleHolidayItem struct {
	CalendarType string `json:"calendar"`
	HolidayType  string `json:"type"`
	HolidayDate  string `json:"date"`
}
