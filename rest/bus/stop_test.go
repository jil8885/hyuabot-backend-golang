package bus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetBusStopList(t *testing.T) {
	test := assert.New(t)
	// Get all bus routes
	t.Log("TestGetBusStopList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/bus/stop", GetBusStopList)
	request := httptest.NewRequest("GET", "/rest/bus/stop", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	test.Nil(err)

	var obj bus.StopListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)

	test.IsType([]bus.StopListItem{}, obj.Stop)
	test.Greater(len(obj.Stop), 0, "There should be at least one bus route")
	for _, stop := range obj.Stop {
		test.IsType(bus.StopListItem{}, stop)
		test.IsType("", stop.Name)
		test.IsType(0, stop.ID)
	}
	// Get bus routes by name
	t.Log("TestGetBusStopListByName")
	searchKeywords := []string{"한양대ERICA컨벤션센터", "한양대정문", "한양대입구"}
	for _, keyword := range searchKeywords {
		request = httptest.NewRequest("GET", "/rest/bus/stop?name="+keyword, nil)
		res, err = app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		body, err = io.ReadAll(res.Body)
		test.Nil(err)
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType([]bus.StopListItem{}, obj.Stop)
		test.Greater(len(obj.Stop), 0, "There should be at least one bus route")
		for _, stop := range obj.Stop {
			test.IsType(bus.StopListItem{}, stop)
			test.IsType("", stop.Name)
			test.IsType(0, stop.ID)
		}
	}
}

func TestGetBusStopItem(t *testing.T) {
	test := assert.New(t)
	// Get all bus routes
	t.Log("TestGetBusStopItem")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/bus/stop/:stop_id", GetBusStopItem)
	stopList := []int{216000070, 216000117, 216000152, 216000138, 216000379, 216000759, 216000719}
	for _, stopID := range stopList {
		request := httptest.NewRequest("GET", fmt.Sprintf("/rest/bus/stop/%d", stopID), nil)
		res, err := app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		test.Nil(err)

		var obj bus.StopItemResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType(0, obj.ID)
		test.IsType("", obj.Name)
		test.IsType("", obj.MobileNumber)
		test.Equal(5, len(strings.Trim(obj.MobileNumber, " ")))
		test.IsType(0.0, obj.Location.Latitude)
		test.IsType(0.0, obj.Location.Longitude)
		test.IsType([]bus.StopRouteItem{}, obj.Route)
		test.Greater(len(obj.Route), 0, "There should be at least one bus route")
		for _, route := range obj.Route {
			test.IsType(bus.StopRouteItem{}, route)
			test.IsType(0, route.ID)
			test.IsType("", route.Name)
			test.IsType(bus.StartStop{}, route.Start)
			test.IsType(0, route.Start.StopID)
			test.IsType("", route.Start.StopName)
			test.IsType([]string{}, route.Start.TimetableList)
			test.IsType([]bus.RealtimeItem{}, route.RealtimeList)
			for _, realtime := range route.RealtimeList {
				test.IsType(bus.RealtimeItem{}, realtime)
				test.IsType(0, realtime.RemainingTime)
				test.IsType(0, realtime.RemainingStop)
				test.IsType(true, realtime.LowPlate)
				test.IsType("", realtime.LastUpdate)
			}
		}
	}
}
