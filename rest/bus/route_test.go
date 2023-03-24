package bus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetBusRouteList(t *testing.T) {
	test := assert.New(t)
	// Get all bus routes
	t.Log("TestGetBusRouteList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/bus/route", GetBusRouteList)
	request := httptest.NewRequest("GET", "/rest/bus/route", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	test.Nil(err)

	var obj bus.RouteListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)

	test.IsType([]bus.RouteListItem{}, obj.Route)
	test.Greater(len(obj.Route), 0, "There should be at least one bus route")
	for _, route := range obj.Route {
		test.IsType(bus.RouteListItem{}, route)
		test.IsType("", route.Name)
		test.IsType(0, route.ID)
	}

	// Get bus routes by name
	t.Log("TestGetBusRouteListByName")
	searchKeywords := []string{"10", "110", "707", "909", "3100", "3101", "3102"}
	for _, keyword := range searchKeywords {
		request = httptest.NewRequest("GET", "/rest/bus/route?name="+keyword, nil)
		res, err = app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)

		body, err = io.ReadAll(res.Body)
		test.Nil(err)

		err = json.Unmarshal(body, &obj)
		test.Nil(err)

		test.IsType([]bus.RouteListItem{}, obj.Route)
		test.Greater(len(obj.Route), 0, "There should be at least one bus route")
		for _, route := range obj.Route {
			test.IsType(bus.RouteListItem{}, route)
			test.IsType("", route.Name)
			test.IsType(0, route.ID)
			test.Contains(route.Name, keyword, "The route name should contain the keyword")
		}
	}
}

func TestGetBusRouteItem(t *testing.T) {
	test := assert.New(t)
	// Get bus route item
	t.Log("TestGetBusRouteItem")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/bus/route/:route_id", GetBusRouteItem)

	routeList := []int{216000061, 216000001, 216000068, 216000096, 216000070, 217000014, 216000016,
		200000015, 216000026, 216000075, 216000043}
	for _, routeID := range routeList {
		fmt.Println("Testing route", routeID)
		request := httptest.NewRequest("GET", fmt.Sprintf("/rest/bus/route/%d", routeID), nil)
		res, err := app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)

		body, err := io.ReadAll(res.Body)
		test.Nil(err)

		var obj bus.RouteItemResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)

		test.IsType(bus.RouteItemResponse{}, obj)
		test.IsType("", obj.Name)
		test.IsType(0, obj.ID)
		test.IsType(bus.RouteType{}, obj.Type)
		test.IsType(bus.RouteCompany{}, obj.Company)
		test.IsType(bus.RouteRunningTimeGroup{}, obj.RunningTime)
		test.IsType(bus.RouteStop{}, obj.Start)
		test.IsType(bus.RouteStop{}, obj.End)

		test.IsType("", obj.Type.Name)
		test.IsType(0, obj.Type.ID)
		test.IsType("", obj.Company.Name)
		test.IsType(0, obj.Company.ID)
		test.IsType("", obj.RunningTime.Up.FirstTime)
		test.IsType("", obj.RunningTime.Up.LastTime)
		test.IsType("", obj.RunningTime.Down.FirstTime)
		test.IsType("", obj.RunningTime.Down.LastTime)
		test.IsType("", obj.Start.Name)
		test.IsType(0, obj.Start.ID)
		test.IsType("", obj.End.Name)
		test.IsType(0, obj.End.ID)
	}
}
