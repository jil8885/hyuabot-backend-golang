package bus

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRealtimeBusDeparture(stopID string, busID string) []DepartureItem {
	authKey := os.Getenv("bus_auth")
	url := "http://apis.data.go.kr/6410000/busarrivalservice/getBusArrivalItem?serviceKey=" + authKey  + "&stationId=" + stopID + "&routeId=" + busID
	// API 서버 데이터 요청
	response, err := http.Get(url)
	result := make([]DepartureItem, 0)
	xmlObj := Response{}

	if err != nil || response.StatusCode != 200 {
		return result
	} else {
		if response.Body != nil {
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(response.Body)
		}
		body, _ := ioutil.ReadAll(response.Body)
		_ = xml.Unmarshal(body, &xmlObj)

	}

	if len(xmlObj.MsgBody.BusArrivalList) > 0{
		item := xmlObj.MsgBody.BusArrivalList[0]
		result = append(result, DepartureItem{Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1})
		if item.LocationNo2 != 0{
			result = append(result, DepartureItem{Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2})
		}
		print(result)
	}
	return result
}

func GetRealtimeStopDeparture(stopID string) StopInfo {
	authKey := os.Getenv("bus_auth")
	url := "http://apis.data.go.kr/6410000/busarrivalservice/getBusArrivalList?serviceKey=" + authKey  + "&stationId=" + stopID
	// API 서버 데이터 요청
	response, err := http.Get(url)
	result := StopInfo{}
	xmlObj := Response{}

	if err != nil || response.StatusCode != 200 {
		return result
	} else {
		if response.Body != nil {
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(response.Body)
		}
		body, _ := ioutil.ReadAll(response.Body)
		_ = xml.Unmarshal(body, &xmlObj)
	}
	fmt.Println(xmlObj)
	return result
}

