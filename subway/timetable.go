package subway

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTimetableSubway() TimetableDataResult {
	var timetableJsonObj TimetableDataByDay
	timetableResult := TimetableDataResult{}

	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	weekdays := getDate(now)
	path, _ := os.Getwd()
	timetableJson := path + "/subway/timetable.json"
	data, err := os.Open(timetableJson)
	if err != nil{
		return TimetableDataResult{}
	}
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &timetableJsonObj)
	if err != nil {
		return TimetableDataResult{}
	}

	if weekdays{
		for _, item := range timetableJsonObj.Weekdays.UpLine{
			if compareTimetable(item.Time, now){
				timetableResult.UpLine = append(timetableResult.UpLine, item)
				if len(timetableResult.UpLine) >= 2{
					break
				}
			}
		}

		for _, item := range timetableJsonObj.Weekdays.DownLine{
			if compareTimetable(item.Time, now){
				timetableResult.DownLine = append(timetableResult.DownLine, item)
				if len(timetableResult.DownLine) >= 2{
					break
				}
			}
		}
	} else {
		for _, item := range timetableJsonObj.Weekend.UpLine{
			if compareTimetable(item.Time, now){
				timetableResult.UpLine = append(timetableResult.UpLine, item)
				if len(timetableResult.UpLine) >= 2{
					break
				}
			}
		}

		for _, item := range timetableJsonObj.Weekend.DownLine{
			if compareTimetable(item.Time, now){
				timetableResult.DownLine = append(timetableResult.DownLine, item)
				if len(timetableResult.DownLine) >= 2{
					break
				}
			}
		}
	}
	return timetableResult
}

func compareTimetable(timeString string, now time.Time) bool {
	slice := strings.Split(timeString, ":")
	hour, _ := strconv.Atoi(slice[0])
	minute, _ := strconv.Atoi(slice[1])
	second, _ := strconv.Atoi(slice[2])

	if hour > now.Hour(){
		return true
	} else if hour == now.Hour(){
		if minute > now.Minute(){
			return true
		} else if minute == now.Minute() && second >= now.Second(){
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}