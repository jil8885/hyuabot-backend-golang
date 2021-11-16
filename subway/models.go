package subway

import "time"

// 전철 도착 정보
type RealtimeDataItem struct {
	UpdatedTime time.Time `firestore:"updatedTime"`
	TerminalStation string `firestore:"terminalStation"`
	Position string `firestore:"position"`
	RemainedTime float64 `firestore:"remainedTime"`
	Status string `firestore:"status"`
}


// 전철 도착 결과
type RealtimeDataResult struct {
	UpLine []RealtimeDataItem `json:"up"`
	DownLine []RealtimeDataItem `json:"down"`
}

// 전철 도착 캐시
type RealtimeDataResultCache struct {
	Result RealtimeAPIResult
	Time time.Time
}

// 전철 API JSON
type RealtimeAPIResult struct {
	ErrorMessage RealtimeAPIError `json:"errorMessage"`
	RealtimeArrivalList []RealtimeAPIItem `json:"realtimeArrivalList"`
}

type RealtimeAPIError struct {
	Status int
}

type RealtimeAPIItem struct {
	UpdatedTime string `json:"recptnDt"`
	LineID string `json:"subwayId"`
	UpDown string `json:"updnLine"`
	TerminalStation string `json:"bstatnNm"`
	CurrentStation string `json:"arvlMsg3"`
	RemainedTime string `json:"barvlDt"`
	Status string `json:"arvlCd"`
}

// 전철 도착 정보
type TimetableDataItem struct {
	TerminalStation string `json:"endStn"`
	Time string `json:"time"`
}


// 전철 도착 결과
type TimetableDataResult struct {
	UpLine []TimetableDataItem `json:"up"`
	DownLine []TimetableDataItem `json:"down"`
}

type TimetableDataByDay struct {
	Weekdays TimetableDataResult `json:"weekdays"`
	Weekend TimetableDataResult `json:"weekend"`
}