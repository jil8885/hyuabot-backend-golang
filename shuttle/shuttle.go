package shuttle

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetShuttle(busStop string, now time.Time) ([]Departure, []Departure) {
	category1, category2 := getDate(now)
	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/" + category1 + "/" + category2 + "/" + busStop + "_" + category2 + ".json"

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

	for _, item := range departureList{
		if strings.Compare(item.Time, strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())) > 0{
			if busStop == "Shuttlecock_I" || busStop == "Terminal" {
				busForTerminal = append(busForTerminal, item)
				if len(busForTerminal) >= 2{
					break
				}
			} else {
				if item.Heading == "C" || item.Heading == ""{
					if len(busForTerminal) < 2{
						busForTerminal = append(busForTerminal, item)
					}

					if len(busForStation) < 2{
						busForStation = append(busForStation, item)
					}
				} else if item.Heading == "DH" && len(busForStation) < 2{
					busForStation = append(busForStation, item)
				}
				if len(busForTerminal) >= 2 && len(busForStation) >= 2{
					break
				}
			}
		}
	}

	return busForStation, busForTerminal
}


func GetFirstLastShuttle(busStop string, now time.Time) (Departure, Departure, Departure, Departure) {
	category1, category2 := getDate(now)
	path, _ := os.Getwd()
	dateJson := path + "/shuttle/timetable/" + category1 + "/" + category2 + "/" + busStop + "_" + category2 + ".json"

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

	for _, item := range departureList{
		if busStop == "Shuttlecock_I" || busStop == "Terminal" {
			busForTerminal = append(busForTerminal, item)
		} else {
			if item.Heading == "C" || item.Heading == ""{
				busForTerminal = append(busForTerminal, item)
				busForStation = append(busForStation, item)
			} else if item.Heading == "DH"{
				busForStation = append(busForStation, item)
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
