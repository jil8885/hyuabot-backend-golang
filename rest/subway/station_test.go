package subway

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/response/subway"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetStationList(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetStationList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/subway/station", GetStationList)
	request := httptest.NewRequest("GET", "/rest/subway/station", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	test.Nil(err)
	var obj subway.StationListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)
	test.IsType([]subway.StationListItem{}, obj.Station)
	test.Greater(len(obj.Station), 0, "There should be at least one subway station")
	for _, station := range obj.Station {
		test.IsType(subway.StationListItem{}, station)
		test.IsType("", station.StationName)
		test.IsType("", station.StationID)
		test.IsType(0, station.RouteID)
	}

	request = httptest.NewRequest("GET", "/rest/subway/station?name=한대앞", nil)
	res, err = app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err = io.ReadAll(res.Body)
	test.Nil(err)
	err = json.Unmarshal(body, &obj)
	test.Nil(err)
	test.IsType([]subway.StationListItem{}, obj.Station)
	test.Greater(len(obj.Station), 0, "There should be at least one subway station")
	for _, station := range obj.Station {
		test.IsType(subway.StationListItem{}, station)
		test.IsType("", station.StationName)
		test.IsType("", station.StationID)
		test.IsType(0, station.RouteID)
	}
}
