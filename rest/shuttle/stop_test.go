package shuttle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetShuttleStopList(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetShuttleStopList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/shuttle/stop", GetShuttleStopList)
	request := httptest.NewRequest("GET", "/rest/shuttle/stop", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	test.Nil(err)
	var obj shuttle.StopListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)
	test.IsType([]shuttle.StopListItem{}, obj.Stop)
	test.Greater(len(obj.Stop), 0, "There should be at least one shuttle route")
	for _, stop := range obj.Stop {
		test.IsType(shuttle.StopListItem{}, stop)
		test.IsType("", stop.Name)
		test.IsType(shuttle.StopLocation{}, stop.Location)
		test.IsType(0.0, stop.Location.Latitude)
		test.IsType(0.0, stop.Location.Longitude)
	}
}

func TestGetShuttleStopItem(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetShuttleStopItem")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/shuttle/stop/:stop_id", GetShuttleStopItem)
	stopList := []string{"dormitory_o", "shuttlecock_o", "station", "terminal", "jungang_stn", "dormitory_i", "shuttlecock_i"}
	for _, stop := range stopList {
		request := httptest.NewRequest("GET", fmt.Sprintf("/rest/shuttle/stop/%s", stop), nil)
		res, err := app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		test.Nil(err)
		var obj shuttle.StopItemResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType(shuttle.StopItemResponse{}, obj)
		test.Equal(stop, obj.Name)
		test.IsType(shuttle.StopLocation{}, obj.Location)
		test.IsType(0.0, obj.Location.Latitude)
		test.IsType(0.0, obj.Location.Longitude)
		test.IsType([]shuttle.StopRouteItem{}, obj.RouteList)
		test.Greater(len(obj.RouteList), 0, "There should be at least one shuttle route")
		for _, route := range obj.RouteList {
			test.IsType(shuttle.StopRouteItem{}, route)
			test.IsType("", route.Name)
			test.IsType("", route.Tag)
			test.IsType(shuttle.RunningTime{}, route.RunningTime)
			test.IsType(shuttle.FirstLastTime{}, route.RunningTime.Weekdays)
			test.IsType("", route.RunningTime.Weekdays.FirstTime)
			test.IsType("", route.RunningTime.Weekdays.LastTime)
			test.IsType(shuttle.FirstLastTime{}, route.RunningTime.Weekends)
			test.IsType("", route.RunningTime.Weekends.FirstTime)
			test.IsType("", route.RunningTime.Weekends.LastTime)
			test.IsType([]string{}, route.TimetableList)
		}
	}
}
