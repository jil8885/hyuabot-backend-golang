package subway

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)


func GetRealtimeSubway(campus int) RealtimeDataResult {
	minuteToArrival := map[string]float64{
		"한대앞": 0, "중앙": 2, "고잔": 4, "초지": 6.5, "안산": 9, "신길온천": 12.5, "정왕": 16, "오이도": 19, "달월": 21, "월곶": 23,
			"소래포구": 25, "인천논현": 27, "호구포": 29, "상록수": 2, "반월": 6, "대야미": 8.5, "수리산": 11.5, "산본": 13.5, "금정": 18,
			"범계": 21.5, "평촌": 23.5, "인덕원": 26, "정부과천청사": 28, "과천": 30, "사리": 2, "야목": 7, "어천": 10, "오목천": 14,
			"고색": 17, "수원": 21, "매교": 23, "수원시청": 26, "매탄권선": 29}

	statusCode := map[int]string{0: "진입", 1: "도착", 2: "출발", 3: "전역출발", 4: "전역진입", 5: "전역도착", 99: "운행중"}

	authKey := os.Getenv("metro_auth")
	url := "http://swopenapi.seoul.go.kr/api/subway/" + strings.TrimSpace(authKey) + "/json/realtimeStationArrival/0/10/"
	if campus == 1 {
		url += "한양대"
	} else {
		url += "한대앞"
	}

	// API 서버 데이터 요청
	result := RealtimeDataResult{}
	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return result
	}
	if response.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)
	}
	body, err := ioutil.ReadAll(response.Body)
	var apiResult RealtimeAPIResult
	var remainedTime float64
	var status int
	err = json.Unmarshal(body, &apiResult)
	// API json 결과 분리
	for _, item := range apiResult.RealtimeArrivalList{
		if campus == 1{
			remainedTime, _ = strconv.ParseFloat(item.RemainedTime, 32)
		} else {
			remainedTime = minuteToArrival[item.CurrentStation]
		}
		status, _ = strconv.Atoi(item.Status)

		if !strings.Contains(item.TerminalStation, "급행"){
			if strings.Contains(item.UpDown, "상행") || strings.Contains(item.UpDown, "내선"){
				result.UpLine = append(result.UpLine, RealtimeDataItem{item.TerminalStation, item.CurrentStation, remainedTime, statusCode[status]})
			} else{
				result.DownLine = append(result.DownLine, RealtimeDataItem{item.TerminalStation, item.CurrentStation, remainedTime, statusCode[status]})
			}
		}
	}
	return result
}
