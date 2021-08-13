package bus

import (
	"encoding/xml"
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
	xmlObj := RouteResponse{}
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

	item := xmlObj.MsgBody.BusArrivalItem
	result = append(result, DepartureItem{Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1})
	if item.LocationNo2 != 0{
		result = append(result, DepartureItem{Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2})
	}
	return result
}

func GetRealtimeStopDeparture(stopID string) StopResponse {
	authKey := os.Getenv("bus_auth")
	url := "http://apis.data.go.kr/6410000/busarrivalservice/getBusArrivalList?serviceKey=" + authKey  + "&stationId=" + stopID
	// API 서버 데이터 요청
	response, err := http.Get(url)
	xmlObj := StopResponse{}

	if err != nil || response.StatusCode != 200 {
		return xmlObj
	} else {
		if response.Body != nil {
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(response.Body)
		}
		body, _ := ioutil.ReadAll(response.Body)
		_ = xml.Unmarshal(body, &xmlObj)
	}
	return xmlObj
}

