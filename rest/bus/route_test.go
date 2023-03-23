package bus

import (
	"encoding/json"
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
