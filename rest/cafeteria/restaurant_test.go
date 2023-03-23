package cafeteria

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/cafeteria"

	"github.com/golang-module/carbon/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetRestaurantList(t *testing.T) {
	test := assert.New(t)
	// Get all bus routes
	t.Log("TestGetRestaurantList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/cafeteria/campus/:campus_id/restaurant", GetRestaurantList)
	campusList := []int{0, 1}
	timetype := []string{"조식", "중식", "석식"}
	for _, campus := range campusList {
		for i := -5; i < 5; i++ {
			date := carbon.Now("Asia/Seoul").AddDays(i).Format("Y-m-d")
			for _, time := range timetype {
				request := httptest.NewRequest("GET", fmt.Sprintf("/rest/cafeteria/campus/%d/restaurant?date=%s&type=%s", campus, date, time), nil)
				res, err := app.Test(request)
				test.Nil(err)
				test.Equal(200, res.StatusCode)
				body, err := io.ReadAll(res.Body)
				test.Nil(err)
				var obj cafeteria.RestaurantListResponse
				err = json.Unmarshal(body, &obj)
				test.Nil(err)
				test.IsType([]cafeteria.RestaurantItemResponse{}, obj.RestaurantList)
				for _, restaurant := range obj.RestaurantList {
					test.IsType(cafeteria.RestaurantItemResponse{}, restaurant)
					test.IsType(0, restaurant.RestaurantID)
					test.IsType("", restaurant.Name)
					test.IsType([]cafeteria.MenuList{}, restaurant.MenuList)
					for _, menu := range restaurant.MenuList {
						test.IsType(cafeteria.MenuList{}, menu)
						test.IsType("", menu.TimeType)
						test.IsType([]cafeteria.MenuItem{}, menu.MenuList)
						for _, menuItem := range menu.MenuList {
							test.IsType(cafeteria.MenuItem{}, menuItem)
							test.IsType("", menuItem.Food)
							test.IsType("0", menuItem.Price)
						}
					}
				}
			}
		}
	}
}

func TestGetRestaurantItem(t *testing.T) {
	test := assert.New(t)
	// Get all bus routes
	t.Log("TestGetRestaurantItem")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/cafeteria/campus/:campus_id/restaurant/:restaurant_id", GetRestaurantItem)
	restaurantList := map[int][]int{
		0: {1, 2, 4, 5, 7, 8},
		1: {11, 12, 13, 15},
	}
	for campus, restaurants := range restaurantList {
		for _, restaurant := range restaurants {
			for i := -5; i < 5; i++ {
				date := carbon.Now("Asia/Seoul").AddDays(i).Format("Y-m-d")
				for _, time := range []string{"조식", "중식", "석식"} {
					request := httptest.NewRequest("GET", fmt.Sprintf("/rest/cafeteria/campus/%d/restaurant/%d?date=%s&type=%s", campus, restaurant, date, time), nil)
					res, err := app.Test(request)
					test.Nil(err)
					test.Equal(200, res.StatusCode)
					body, err := io.ReadAll(res.Body)
					test.Nil(err)
					var obj cafeteria.RestaurantListResponse
					err = json.Unmarshal(body, &obj)
					test.Nil(err)
					test.IsType([]cafeteria.RestaurantItemResponse{}, obj.RestaurantList)
					for _, restaurant := range obj.RestaurantList {
						test.IsType(cafeteria.RestaurantItemResponse{}, restaurant)
						test.IsType(0, restaurant.RestaurantID)
						test.IsType("", restaurant.Name)
						test.IsType([]cafeteria.MenuList{}, restaurant.MenuList)
						for _, menu := range restaurant.MenuList {
							test.IsType(cafeteria.MenuList{}, menu)
							test.IsType("", menu.TimeType)
							test.IsType([]cafeteria.MenuItem{}, menu.MenuList)
							for _, menuItem := range menu.MenuList {
								test.IsType(cafeteria.MenuItem{}, menuItem)
								test.IsType("", menuItem.Food)
								test.IsType("0", menuItem.Price)
							}
						}
					}
				}
			}
		}
	}
}
