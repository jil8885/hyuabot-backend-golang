package shuttle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetShuttle(busStop string, now time.Time, loc *time.Location) ([]Departure, []Departure) {
	category1, category2 := GetDate(now, loc)
	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/" + category1 + "/" + category2 + ".json"

	// 시간표 json 로딩
	var departureList []Departure
	data, err := os.Open(dateJson)
	if err != nil{
		return []Departure{}, []Departure{}
	}
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &departureList)
	if err != nil {
		return []Departure{}, []Departure{}
	}

	// 반환할 데이터 선택
	var busForStation []Departure
	var busForTerminal []Departure
	var timedelta = 0

	for _, item := range departureList{
		if busStop == "Residence"{
			timedelta = -15
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_O"{
			timedelta = -10
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Subway"{
			timedelta = 0
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Terminal"{
			timedelta = 5
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DY" || item.Heading == "C"{
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_I"{
			if item.Heading == "DH" || item.Heading == "DY"{
				timedelta = 10
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			} else if item.Heading == "C"{
				timedelta = 15
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			}
		}


		if len(busForStation) >= 2 && len(busForTerminal) >= 2{
			break
		}
	}

	if busForStation == nil {
		busForStation = []Departure{}
	} else if len(busForStation) >= 2{
		busForStation = busForStation[0:2]
	}

	if busForTerminal == nil {
		busForTerminal = []Departure{}
	} else if len(busForTerminal) >= 2{
		busForTerminal = busForTerminal[0:2]
	}
	return busForStation, busForTerminal
}

func GetShuttleTimetable(busStop string, now time.Time, loc *time.Location, category string) ([]Departure, []Departure) {
	category1, _ := GetDate(now, loc)
	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/" + category1 + "/" + category + ".json"

	// 시간표 json 로딩
	var departureList []Departure
	data, err := os.Open(dateJson)
	if err != nil{
		return []Departure{}, []Departure{}
	}
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &departureList)
	if err != nil {
		return []Departure{}, []Departure{}
	}

	// 반환할 데이터 선택
	var busForStation []Departure
	var busForTerminal []Departure

	var timedelta = 0
	for _, item := range departureList{
		if busStop == "Residence"{
			timedelta = -15
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_O"{
			timedelta = -10
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Subway"{
			timedelta = 0
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Terminal"{
			timedelta = 5
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DY" || item.Heading == "C"{
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_I"{
			if item.Heading == "DH" || item.Heading == "DY"{
				timedelta = 10
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			} else if item.Heading == "C"{
				timedelta = 15
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			}
		}
	}

	return busForStation, busForTerminal
}

func GetFirstLastShuttle(busStop string, now time.Time, loc *time.Location) (Departure, Departure, Departure, Departure) {
	category1, category2 := GetDate(now, loc)
	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/" + category1 + "/" + category2 + ".json"

	// 시간표 json 로딩
	var departureList []Departure
	data, err := os.Open(dateJson)
	if err != nil{
		return Departure{}, Departure{}, Departure{}, Departure{}
	}
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &departureList)
	if err != nil {
		return Departure{}, Departure{}, Departure{}, Departure{}
	}

	// 반환할 데이터 선택
	var busForStation []Departure
	var busForTerminal []Departure

	var timedelta = 0
	for _, item := range departureList{
		if busStop == "Residence"{
			timedelta = -15
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_O"{
			timedelta = -10
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "DY"{
					busForTerminal = append(busForTerminal, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Subway"{
			timedelta = 0
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DH"{
					busForStation = append(busForStation, item)
				} else if item.Heading == "C"{
					busForStation = append(busForStation, item)
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Terminal"{
			timedelta = 5
			if compareTimetable(item.Time, now, timedelta){
				item.Time = getTimeFromTimeDelta(item.Time, timedelta)
				if item.Heading == "DY" || item.Heading == "C"{
					busForTerminal = append(busForTerminal, item)
				}
			}
		} else if busStop == "Shuttlecock_I"{
			if item.Heading == "DH" || item.Heading == "DY"{
				timedelta = 10
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			} else if item.Heading == "C"{
				timedelta = 15
				if compareTimetable(item.Time, now, timedelta){
					item.Time = getTimeFromTimeDelta(item.Time, timedelta)
					if item.Heading == "DY" || item.Heading == "C"{
						busForTerminal = append(busForTerminal, item)
					}
				}
			}
		}
	}
	if busForStation != nil && busForTerminal != nil{
		return busForStation[0], busForStation[len(busForStation) - 1], busForTerminal[0], busForTerminal[len(busForTerminal) - 1]

	} else if busForStation != nil{
		return busForStation[0], busForStation[len(busForStation) - 1], Departure{}, Departure{}
	} else if busForTerminal != nil{
		return Departure{}, Departure{}, busForTerminal[0], busForTerminal[len(busForTerminal) - 1]
	} else{
		return Departure{}, Departure{}, Departure{}, Departure{}
	}
}

func compareTimetable(timeString string, now time.Time, timedelta int) bool {
	slice := strings.Split(timeString, ":")
	hour, _ := strconv.Atoi(slice[0])
	minute, _ := strconv.Atoi(slice[1])
	minute += timedelta
	if minute < 0{
		hour -= 1
		minute += 60
	} else if minute >= 60{
		hour +=1
		minute -= 60
	}

	if hour > now.Hour(){
		return true
	} else if hour == now.Hour(){
		if minute > now.Minute(){
			return true
		}else {
			return false
		}
	} else {
		return false
	}
}

func getTimeFromTimeDelta(timeString string, timedelta int) string {
	slice := strings.Split(timeString, ":")
	hour, _ := strconv.Atoi(slice[0])
	minute, _ := strconv.Atoi(slice[1])
	minute += timedelta
	if minute < 0{
		hour -= 1
		minute += 60
	} else if minute >= 60{
		hour +=1
		minute -= 60
	}

	return fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", minute)
}