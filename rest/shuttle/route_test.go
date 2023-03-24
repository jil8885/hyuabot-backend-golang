package shuttle

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetShuttleRouteList(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetShuttleRouteList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/shuttle/route", GetShuttleRouteList)
	request := httptest.NewRequest("GET", "/rest/shuttle/route", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	test.Nil(err)
	var obj shuttle.RouteListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)
	test.IsType([]shuttle.RouteListItem{}, obj.Route)
	test.Greater(len(obj.Route), 0, "There should be at least one shuttle route")
	for _, route := range obj.Route {
		test.IsType(shuttle.RouteListItem{}, route)
		test.IsType("", route.Name)
		test.IsType("", route.Tag)
		test.IsType(shuttle.RouteDescription{}, route.Description)
		test.IsType("", route.Description.Korean)
		test.IsType("", route.Description.English)
	}
}
