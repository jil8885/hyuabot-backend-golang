package library

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"net/http"
)

func FetchLibrary()  {
	url := "https://lib.hanyang.ac.kr/smufu-api/pc/2/rooms-at-seat"

	// API 서버 데이터 요청
	result := ReadingRoomJSON{}
	response, err := http.Get(url)

	roomMap := map[string]string{"제1열람실": "reading_room_1", "제2열람실": "reading_room_2", "제3열람실": "reading_room_3", "제4열람실": "reading_room_4"}

	if err != nil{
		return
	}
	if response.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(response.Body)
	}

	body, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(body, &result)

	fmt.Println(result)

	if !result.Success{
		fmt.Println("Result failed")
		return
	}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil{
		fmt.Println(err)
		return
	}

	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		return
	}
	// Firestore handling
	err = client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		document := client.Collection("hanyangApp").Doc("readingRoom")
		_, err = document.Set(ctx, map[string]interface{}{
			"updateTime": firestore.ServerTimestamp,
		}, firestore.MergeAll)

		collection := document.Collection("rooms")
		for _, item := range  result.Data.ReadingRoomList{
			_, err = collection.Doc(item.Name).Set(ctx, map[string]interface{}{
				"isActive": item.IsActive,
				"IsReservable": item.IsReservable,
				"total": item.Total,
				"activeTotal": item.ActiveTotal,
				"occupied": item.Occupied,
				"available": item.Available,
			}, firestore.MergeAll)
			
			if item.IsReservable && item.Available > 0{
				message := &messaging.Message{
					Data:         map[string]string{"type": "reading_room", "name": roomMap[item.Name]},
					Android:      &messaging.AndroidConfig{Priority: "high"},
					Topic:        roomMap[item.Name],
				}

				_, err := messagingClient.Send(ctx, message)
				if err != nil {
					return err
				}
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
	}(client)
	return
}

func GetLibrary() []ReadingRoomInfo {
	var queryResult []ReadingRoomInfo

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		return nil
	}

	client, err := app.Firestore(ctx)
	if err != nil{
		fmt.Println(err)
		return nil
	}

	// Firestore handling
	iter := client.Collection("hanyangApp").Doc("readingRoom").Collection("rooms").Documents(ctx)
	for{
		item, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		var room ReadingRoomInfo
		err = item.DataTo(&room)
		if err != nil {
			return nil
		}
		room.Name = item.Ref.ID
		queryResult = append(queryResult, room)
	}


	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(client)
	return queryResult
}

func GetLibraryByName(name string) ReadingRoomInfo {
	var queryResult ReadingRoomInfo

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		return queryResult
	}

	client, err := app.Firestore(ctx)
	if err != nil{
		fmt.Println(err)
		return queryResult
	}

	// Firestore handling
	item, err := client.Collection("hanyangApp").Doc("readingRoom").Collection("rooms").Doc(name).Get(ctx)
	if err != nil {
		return queryResult
	}

	err = item.DataTo(&queryResult)
	queryResult.Name = item.Ref.ID


	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(client)
	return queryResult
}