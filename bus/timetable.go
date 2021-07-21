package bus

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetTimetable() BusTimeTableJson {
	path, _ := os.Getwd()
	timetableJson := path + "/bus/timetable.json"

	var timetableObj BusTimeTableJson
	data, err := os.Open(timetableJson)
	if err != nil{
		return BusTimeTableJson{}
	}

	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &timetableObj)

	if err != nil{
		return BusTimeTableJson{}
	} else {
		return timetableObj
	}

}