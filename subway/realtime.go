package subway

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetRealtimeSubway(campus int, lineID int) RealtimeDataResult {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)

	loc, _ := time.LoadLocation("Asia/Seoul")

	if err != nil {
		return RealtimeDataResult{}
	}

	client, err := app.Firestore(ctx)
	if err != nil{
		fmt.Println(err)
		return RealtimeDataResult{}
	}

	// Firestore handling
	var result RealtimeDataResult
	iter := client.Collection("hanyangApp").Doc("subway").Collection("campus").Doc(strconv.Itoa(campus)).Collection("lineID").Doc(strconv.Itoa(lineID)).Collection("upLine").Documents(ctx)
	for{
		item, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
		}
		var realtimeDataItem RealtimeDataItem
		err = item.DataTo(&realtimeDataItem)
		if err != nil {
			continue
		}
		realtimeDataItem.UpdatedTime = realtimeDataItem.UpdatedTime.In(loc)
		result.UpLine = append(result.UpLine, realtimeDataItem)
	}

	iter = client.Collection("hanyangApp").Doc("subway").Collection("campus").Doc(strconv.Itoa(campus)).Collection("lineID").Doc(strconv.Itoa(lineID)).Collection("downLine").Documents(ctx)
	for{
		item, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
		}
		var realtimeDataItem RealtimeDataItem
		err = item.DataTo(&realtimeDataItem)
		if err != nil {
			continue
		}
		realtimeDataItem.UpdatedTime = realtimeDataItem.UpdatedTime.In(loc)
		result.DownLine = append(result.DownLine, realtimeDataItem)
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(client)
	return result
}

func FetchSubwayRealtime(campus int, lineID int)  {
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
	result := RealtimeDataResult{[]RealtimeDataItem{}, []RealtimeDataItem{}}
	client := http.Client{Timeout: 3*time.Second}

	response, err := client.Get(url)
	var apiResult RealtimeAPIResult
	var remainedTime float64
	var status int

	if response.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(response.Body)
	}
	body, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(body, &apiResult)

	// API json 결과 분리
	for _, item := range apiResult.RealtimeArrivalList{
		if campus == 1{
			remainedTime, _ = strconv.ParseFloat(item.RemainedTime, 32)
		} else {
			remainedTime = minuteToArrival[item.CurrentStation]
		}
		status, _ = strconv.Atoi(item.Status)
		loc, _ := time.LoadLocation("Asia/Seoul")

		if !strings.Contains(item.TerminalStation, "급행") && item.LineID == strconv.Itoa(lineID){
			if strings.Contains(item.UpDown, "상행") || strings.Contains(item.UpDown, "내선"){
				updateTime, _ := time.ParseInLocation("2006-01-02 15:04:05.0", item.UpdatedTime, loc)
				result.UpLine = append(result.UpLine, RealtimeDataItem{updateTime, item.TerminalStation, item.CurrentStation, remainedTime, statusCode[status]})
			} else{
				updateTime, _ := time.ParseInLocation("2006-01-02 15:04:05.0", item.UpdatedTime, loc)
				result.DownLine = append(result.DownLine, RealtimeDataItem{updateTime,item.TerminalStation, item.CurrentStation, remainedTime, statusCode[status]})
			}
		}
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	firestoreClient, firestoreError := app.Firestore(ctx)
	if firestoreError != nil{
		fmt.Println(firestoreError)
		return
	}

	// Firestore handling
	err = firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		document := firestoreClient.Collection("hanyangApp").Doc("subway").Collection("campus").Doc(strconv.Itoa(campus)).Collection("lineID").Doc(strconv.Itoa(lineID))
		_, err = document.Set(ctx, map[string]interface{}{
			"updateTime": firestore.ServerTimestamp,
		}, firestore.MergeAll)


		collection := document.Collection("upLine")
		for i:=0; i<2; i++ {
			if i >= len(result.UpLine){
				_, err := collection.Doc(strconv.Itoa(i)).Delete(ctx)
				if err != nil {
					return err
				}
			} else {
					_, err = collection.Doc(strconv.Itoa(i)).Set(ctx, map[string]interface{}{
					"updatedTime": result.UpLine[i].UpdatedTime,
					"terminalStation": result.UpLine[i].TerminalStation,
					"position": result.UpLine[i].Position,
					"remainedTime": result.UpLine[i].RemainedTime,
					"status": result.UpLine[i].Status,
				}, firestore.MergeAll)
			}
		}

		collection = document.Collection("downLine")
		for i:=0; i<2; i++ {
			if i >= len(result.DownLine){
				_, err := collection.Doc(strconv.Itoa(i)).Delete(ctx)
				if err != nil {
					return err
				}
			} else {
				_, err = collection.Doc(strconv.Itoa(i)).Set(ctx, map[string]interface{}{
					"updatedTime": result.DownLine[i].UpdatedTime,
					"terminalStation": result.DownLine[i].TerminalStation,
					"position": result.DownLine[i].Position,
					"remainedTime": result.DownLine[i].RemainedTime,
					"status": result.DownLine[i].Status,
				}, firestore.MergeAll)
			}
		}
		return err
	})


	if err != nil{
		fmt.Println(err)
		return
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(firestoreClient)
	return
}
