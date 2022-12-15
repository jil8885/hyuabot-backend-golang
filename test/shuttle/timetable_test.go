package shuttle

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/test"
)

func TestShuttleTimetable(t *testing.T) {
	// Request Object
	request := httptest.NewRequest("GET", "/rest/shuttle/timetable", nil)
	app := test.InitApp()

	// Test
	var resp shuttle.StopTimetableListResponse
	response, err := app.Test(request)
	if err != nil {
		t.Error(err)
	} else if response.StatusCode != 200 {
		t.Error("Status code is not 200")
	}
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	if resp.Stop == nil {
		t.Error("Response is nil")
	}
	for _, stop := range resp.Stop {
		if stop.Name == "" {
			t.Error("Stop name is empty")
		}
		if stop.Route == nil {
			t.Error("Route is empty")
		}
		for _, route := range stop.Route {
			if route.Name == "" {
				t.Error("Route name is empty")
			}
			if route.Weekdays == nil {
				t.Error("Timetable is empty")
			}
			for _, timetable := range route.Weekdays {
				if timetable == "" {
					t.Error("Time is empty")
				}
			}
			if route.Weekends == nil {
				t.Error("Timetable is empty")
			}
			for _, timetable := range route.Weekends {
				if timetable == "" {
					t.Error("Time is empty")
				}
			}
		}
	}
}

func TestShuttleArrival(t *testing.T) {
	// Request Object
	request := httptest.NewRequest("GET", "/rest/shuttle/arrival", nil)
	app := test.InitApp()

	// Test
	var resp shuttle.StopArrivalListResponse
	response, err := app.Test(request)
	if err != nil {
		t.Error(err)
	} else if response.StatusCode != 200 {
		t.Error("Status code is not 200")
	}
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	if resp.Stop == nil {
		t.Error("Response is nil")
	}
	for _, stop := range resp.Stop {
		if stop.Name == "" {
			t.Error("Stop name is empty")
		}
		if stop.Route == nil {
			t.Error("Route is empty")
		}
		for _, route := range stop.Route {
			if route.Name == "" {
				t.Error("Route name is empty")
			}
			if route.ArrivalList == nil {
				t.Error("Arrival data is empty")
			}
			for _, arrival := range route.ArrivalList {
				if reflect.TypeOf(arrival).Kind() != reflect.Int64 {
					t.Error("Arrival time is " + reflect.TypeOf(arrival).String())
				}
			}
		}
	}
}
