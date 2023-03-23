package bus

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestGetBusRouteList(t *testing.T) {
	testing := assert.New(t)
	// Get all bus routes
	t.Log("TestGetBusRouteList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/bus/route", GetBusRouteList)
	request := httptest.NewRequest("GET", "/rest/bus/route", nil)
	res, err := app.Test(request)
	testing.Nil(err)
	testing.Equal(200, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	testing.Nil(err)

	var obj bus.RouteListResponse
	err = json.Unmarshal(body, &obj)
	testing.Nil(err)

	testing.IsType([]bus.RouteListItem{}, obj.Route)
	testing.Greater(len(obj.Route), 0, "There should be at least one bus route")
	for _, route := range obj.Route {
		testing.IsType(bus.RouteListItem{}, route)
		testing.IsType("", route.Name)
		testing.IsType(0, route.ID)
	}

	// Get bus routes by name
	t.Log("TestGetBusRouteListByName")
	searchKeywords := []string{"10", "110", "707", "909", "3100", "3101", "3102"}
	for _, keyword := range searchKeywords {
		request = httptest.NewRequest("GET", "/rest/bus/route?name="+keyword, nil)
		res, err = app.Test(request)
		testing.Nil(err)
		testing.Equal(200, res.StatusCode)

		body, err = io.ReadAll(res.Body)
		testing.Nil(err)

		err = json.Unmarshal(body, &obj)
		testing.Nil(err)

		testing.IsType([]bus.RouteListItem{}, obj.Route)
		testing.Greater(len(obj.Route), 0, "There should be at least one bus route")
		for _, route := range obj.Route {
			testing.IsType(bus.RouteListItem{}, route)
			testing.IsType("", route.Name)
			testing.IsType(0, route.ID)
			testing.Contains(route.Name, keyword, "The route name should contain the keyword")
		}
	}
}

func TestGetBusRouteItem(t *testing.T) {
	testing := assert.New(t)
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
		testing.Nil(err)
		testing.Equal(200, res.StatusCode)

		body, err := io.ReadAll(res.Body)
		testing.Nil(err)

		var obj bus.RouteItemResponse
		err = json.Unmarshal(body, &obj)
		testing.Nil(err)

		testing.IsType(bus.RouteItemResponse{}, obj)
		testing.IsType("", obj.Name)
		testing.IsType(0, obj.ID)
		testing.IsType(bus.RouteType{}, obj.Type)
		testing.IsType(bus.RouteCompany{}, obj.Company)
		testing.IsType(bus.RouteRunningTimeGroup{}, obj.RunningTime)
		testing.IsType(bus.RouteStop{}, obj.Start)
		testing.IsType(bus.RouteStop{}, obj.End)

		testing.IsType("", obj.Type.Name)
		testing.IsType(0, obj.Type.ID)
		testing.IsType("", obj.Company.Name)
		testing.IsType(0, obj.Company.ID)
		testing.IsType("", obj.RunningTime.Up.FirstTime)
		testing.IsType("", obj.RunningTime.Up.LastTime)
		testing.IsType("", obj.RunningTime.Down.FirstTime)
		testing.IsType("", obj.RunningTime.Down.LastTime)
		testing.IsType("", obj.Start.Name)
		testing.IsType(0, obj.Start.ID)
		testing.IsType("", obj.End.Name)
		testing.IsType(0, obj.End.ID)
	}
}
