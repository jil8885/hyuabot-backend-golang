package bus

import (
	"encoding/json"
	"io"
	"net/http/httptest"
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
