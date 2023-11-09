package requests

import "github.com/volatiletech/null/v8"

type ShuttleTimetableRequest struct {
	PeriodType    string `json:"period" validate:"required"`
	Weekday       bool   `json:"weekday" validate:"required"`
	Route         string `json:"route" validate:"required"`
	DepartureTime string `json:"time" validate:"required"`
}

type ShuttleTimetableUpdateRequest struct {
	PeriodType    null.String `json:"period"`
	Weekday       null.Bool   `json:"weekday"`
	Route         null.String `json:"route"`
	DepartureTime null.String `json:"time"`
}

type ShuttleRouteRequest struct {
	Name               string `json:"name" validate:"required"`
	Tag                string `json:"tag" validate:"required"`
	DescriptionKorean  string `json:"descriptionKorean" validate:"required"`
	DescriptionEnglish string `json:"descriptionEnglish" validate:"required"`
	Start              string `json:"start" validate:"required"`
	End                string `json:"end" validate:"required"`
}

type ShuttleRouteUpdateRequest struct {
	Tag                null.String `json:"tag"`
	DescriptionKorean  null.String `json:"descriptionKorean"`
	DescriptionEnglish null.String `json:"descriptionEnglish"`
	Start              null.String `json:"start"`
	End                null.String `json:"end"`
}

type ShuttleStopRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type ShuttleStopUpdateRequest struct {
	Latitude  null.Float64 `json:"latitude"`
	Longitude null.Float64 `json:"longitude"`
}

type ShuttleRouteStopRequest struct {
	Name           string `json:"name" validate:"required"`
	Seq            int    `json:"seq" validate:"required"`
	CumulativeTime string `json:"cumulativeTime" validate:"required"`
}

type ShuttleRouteStopUpdateRequest struct {
	Seq            null.Int    `json:"seq"`
	CumulativeTime null.String `json:"cumulativeTime"`
}

type ShuttlePeriodRequest struct {
	PeriodType string `json:"type" validate:"required"`
	StartDate  string `json:"start" validate:"required"`
	EndDate    string `json:"end" validate:"required"`
}

type ShuttleHolidayRequest struct {
	CalendarType string `json:"calendar" validate:"required"`
	HolidayType  string `json:"holiday" validate:"required"`
	HolidayDate  string `json:"date" validate:"required"`
}
