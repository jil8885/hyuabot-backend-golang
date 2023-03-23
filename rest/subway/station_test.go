package subway

import (
	"encoding/json"
	"fmt"
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

func TestGetStationItem(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetStationItem")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/subway/station/:station_id", GetStationItem)
	stationList := []string{"K456", "K449", "K258", "K251"}
	for _, stationID := range stationList {
		request := httptest.NewRequest("GET", "/rest/subway/station/"+stationID, nil)
		res, err := app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		test.Nil(err)
		var obj subway.StationItemResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType(subway.StationItemResponse{}, obj)
		test.IsType("", obj.StationName)
		test.IsType("", obj.StationID)
		test.IsType(0, obj.RouteID)
		test.IsType(subway.StationRealtimeHeadingGroup{}, obj.Realtime)
		test.IsType([]subway.StationRealtimeItem{}, obj.Realtime.Up)
		test.IsType([]subway.StationRealtimeItem{}, obj.Realtime.Down)
		for _, item := range obj.Realtime.Up {
			test.IsType(subway.StationRealtimeItem{}, item)
			test.IsType("", item.TerminalStationName)
			test.IsType("", item.TrainNo)
			test.IsType(0, item.RemainingTime)
			test.IsType(0, item.RemainingStop)
			test.IsType(true, item.IsUp)
			test.IsType(true, item.IsExpress)
			test.IsType("", item.CurrentLocation)
			test.IsType("", item.LastUpdate)
		}
		for _, item := range obj.Realtime.Down {
			test.IsType(subway.StationRealtimeItem{}, item)
			test.IsType("", item.TerminalStationName)
			test.IsType("", item.TrainNo)
			test.IsType(0, item.RemainingTime)
			test.IsType(0, item.RemainingStop)
			test.IsType(true, item.IsUp)
			test.IsType(true, item.IsExpress)
			test.IsType("", item.CurrentLocation)
			test.IsType("", item.LastUpdate)
		}

		test.IsType(subway.StationTimetableHeadingGroup{}, obj.Timetable)
		test.IsType([]subway.StationTimetableItem{}, obj.Timetable.Up)
		test.IsType([]subway.StationTimetableItem{}, obj.Timetable.Down)
		for _, item := range obj.Timetable.Up {
			test.IsType(subway.StationTimetableItem{}, item)
			test.IsType("", item.TerminalStationName)
			test.IsType("", item.StartStationName)
			test.IsType("", item.DepartureTime)
		}
	}
}

func TestGetStationArrival(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetStationArrival")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/subway/station/:station_id/arrival", GetStationArrival)
	stationList := []string{"K456", "K449", "K258", "K251"}
	for _, stationID := range stationList {
		request := httptest.NewRequest("GET", fmt.Sprintf("/rest/subway/station/%s/arrival", stationID), nil)
		res, err := app.Test(request)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		test.Nil(err)
		var obj subway.StationArrivalResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType(subway.StationArrivalResponse{}, obj)
		test.IsType("", obj.StationName)
		test.IsType("", obj.StationID)
		test.IsType(0, obj.RouteID)
		test.IsType(subway.StationArrivalHeadingGroup{}, obj.Arrival)
		test.IsType([]subway.StationArrivalItem{}, obj.Arrival.Up)
		test.IsType([]subway.StationArrivalItem{}, obj.Arrival.Down)
		for _, item := range obj.Arrival.Up {
			test.IsType(subway.StationArrivalItem{}, item)
			test.IsType("", item.TerminalStationName)
			test.IsType(0, item.RemainingTime)
			test.IsType(0, item.RemainingStop)
			test.IsType("", item.CurrentStationName)
			test.IsType("", item.DataType)
			test.IsType("", item.LastUpdate)
		}
		for _, item := range obj.Arrival.Down {
			test.IsType(subway.StationArrivalItem{}, item)
			test.IsType("", item.TerminalStationName)
			test.IsType(0, item.RemainingTime)
			test.IsType(0, item.RemainingStop)
			test.IsType("", item.CurrentStationName)
			test.IsType("", item.DataType)
			test.IsType("", item.LastUpdate)
		}
	}
}
