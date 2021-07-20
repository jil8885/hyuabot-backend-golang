package library

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
)

func FetchLibrary()  {
	url := "https://lib.hanyang.ac.kr/smufu-api/pc/2/rooms-at-seat"

	// API 서버 데이터 요청
	result := ReadingRoomJSON{}
	response, err := http.Get(url)

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

func GetLibrary(name string) error {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil{
		fmt.Println(err)
		return nil
	}

	// Firestore handling
	iter := client.Collection("hanyangApp").Doc("readingRoom").Collection("rooms").Documents(ctx)
	for{
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
	}


	if err != nil{
		fmt.Println(err)
		return nil
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(client)
	return nil
}