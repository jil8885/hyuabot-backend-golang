package food

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func FetchFoodMenu(){
	url := "https://www.hanyang.ac.kr/web/www/re"

	// 현재 시간 로딩 (KST)
	//loc, _ := time.LoadLocation("Asia/Seoul")
	//_ := time.Now().In(loc)

	restaurantsMap := map[string]int{"교직원식당": 11, "학생식당": 12, "창의인재원식당": 13, "푸드코트": 14, "창업보육센터": 15}
	var restaurantsList []Restaurant
	for key, value := range restaurantsMap{
		response, err := http.Get(url + strconv.Itoa(value))
		if err != nil{
			continue
		}

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil{
			return
		}

		restaurant := Restaurant{Name: key}
		restaurant.Time = ""
		doc.Find("div.tab-pane").Each(func(i int, s *goquery.Selection) {
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				text := strings.TrimSpace(s.Text())
				if strings.Contains(text, "조식") || strings.Contains(text, "중식") || strings.Contains(text, "석식"){
					restaurant.Time += text + "\n"
				}
			})
		})

		restaurant.MenuList = map[string][]Menu{}
		doc.Find("div.in-box").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h4").First().Text()
			restaurant.MenuList[title] = []Menu{}
			s.Find("li").Each(func(i int, s *goquery.Selection){
				h3 := s.Find("h3")
				p := s.Find("p.price")
				if h3.Length() > 0 && p.Length() > 0{
					menu := strings.TrimSpace(h3.First().Text())
					price, _ := strconv.Atoi(p.First().Text())
					restaurant.MenuList[title] = append(restaurant.MenuList[title], Menu{Menu: menu, Price: price})
				}
			})
		})
		restaurantsList = append(restaurantsList, restaurant)
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
		document := client.Collection("hanyangApp").Doc("food")
		_, err = document.Set(ctx, map[string]interface{}{
			"updateTime": firestore.ServerTimestamp,
		}, firestore.MergeAll)

		collection := document.Collection("restaurants")
		for _, item := range  restaurantsList{
			_, err = collection.Doc(item.Name).Set(ctx, map[string]interface{}{
				"time": item.Time,
				"menuList": item.MenuList,
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
