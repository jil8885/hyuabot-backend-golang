package subway

import (
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"time"
)

func getDate(now time.Time) bool{
	day := true
	if now.Weekday() == 0 || now.Weekday() == 6 || isHoliday(now){
		day = false
	}
	return day
}

func isHoliday(time time.Time) bool {
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
