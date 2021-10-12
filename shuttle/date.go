package shuttle

import (
	"encoding/json"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"io/ioutil"
	"os"
	"time"
)

func GetDate(now time.Time, loc *time.Location) (string, string) {
	// 학기중, 계절학기, 방학 중인지 구별 코드
	// json 파일 로드
	var dateInfo DateJson

	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/date.json"
	data, err := os.Open(dateJson)
	if err != nil{
		return "", ""
	}
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &dateInfo)
	if err != nil {
		return "", ""
	}
	
	// json 파일을 통해 학기/방학/계절학기 구분
	correct := -1
	const layout = "1/2/2006"
	for index, group := range [][]SectionJson{dateInfo.Semester, dateInfo.Vacation, dateInfo.VacationSession}{
		for _, item := range group{
			start, _ := time.Parse(layout, item.StartDate)
			end, _ := time.Parse(layout, item.EndDate)
			yearToADD := now.Year() - start.Year()
			start = start.AddDate(yearToADD, 0, 0).In(loc).Add(-9 * time.Duration(time.Hour))
			end = end.AddDate(yearToADD, 0, 0).In(loc).Add(15 * time.Duration(time.Hour) - 1 * time.Duration(time.Second))
			if now.After(start) && now.Before(end){
				correct = index
				break
			}
		}

		if correct != -1 {
			break
		}
	}

	working := true
	term := ""
	switch correct {
	case 0:
		term = "semester"
	case 1:
		term = "vacation"
	case 2:
		term = "vacation_session"
	default:
		working = false
		term = "vacation"
	}

	if !working{
		for _, item := range dateInfo.Halt{
			date, _ := time.Parse(layout, item)
			if now.Month() == date.Month() && now.Day() == date.Day(){
				term = "halt"
				break
			}
		}
	}

	day := "week"
	if now.Weekday() == 0 || now.Weekday() == 6 || IsHoliday(now){
		day = "weekend"
	}

	for _, holiday := range dateInfo.Holiday{
		date, _ := time.Parse(layout, holiday)
		yearToADD := now.Year() - date.Year()
		date = date.AddDate(yearToADD, 0, 0).In(loc).Add(-9 * time.Duration(time.Hour))
		if now.Year() == date.Year() && now.Month() == date.Month() && now.Day() == date.Day() {
			day = "weekend"
			break
		}
	}
	return term, day
}

func IsHoliday(time time.Time) bool {
	lunarHoliday := []*calendar.Calendar{calendar.ByLunar(int64(time.Year()), 4, 8, 0, 0, 0, false), calendar.ByLunar(int64(time.Year()), 8, 14, 0, 0, 0, false), calendar.ByLunar(int64(time.Year()), 8, 15, 0, 0, 0, false), calendar.ByLunar(int64(time.Year()), 8, 16, 0, 0, 0, false), calendar.ByLunar(int64(time.Year()) - 1, 12, 30, 0, 0, 0, false),calendar.ByLunar(int64(time.Year()), 1, 1, 0, 0, 0, false), calendar.ByLunar(int64(time.Year()), 1, 2, 0, 0, 0, false)}
	boolean := false
	for _, date := range lunarHoliday{
		solarData := date.Solar
		if solarData.GetYear() == int64(time.Year()) && solarData.GetMonth() == int64(time.Month()) && solarData.GetDay() == int64(time.Day()){
			boolean = true
			break
		}
	}
	return boolean
}
