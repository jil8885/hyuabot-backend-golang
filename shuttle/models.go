package shuttle

// 날짜를 담은 json 파일의 모델
type DateJson struct {
	Holiday         []string      `json:"holiday"`
	Halt            []string      `json:"halt"`
	Semester        []SectionJson `json:"semester"`
	VacationSession []SectionJson `json:"vacation_session"`
	Vacation        []SectionJson `json:"vacation"`
}

// 해당 학기, 방학 및 계절학기의 이름, 시작 일자, 종료 일자
type SectionJson struct {
	Key       string `json:"key"`
	StartDate string `json:"start"`
	EndDate   string `json:"end"`
}

type Departure struct {
	Time      string `json:"time"`
	Heading   string `json:"type"`
	StartStop string `json:"startStop"`
}
