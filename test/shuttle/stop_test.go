package shuttle

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/test"
)

func TestShuttleStopList(t *testing.T) {
	// Request Object
	request := httptest.NewRequest("GET", "/rest/shuttle/stop", nil)
	app := test.InitApp()

	// Test
	var resp shuttle.StopListResponse
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
		if stop.Location.Latitude == 0 || stop.Location.Longitude == 0 {
			t.Error("Location is empty")
		}
	}
}

func TestShuttleStopItem(t *testing.T) {
	// Request Object
	stopList := [6]string{"dormitory_o", "shuttlecock_o", "station", "terminal", "shuttlecock_i", "dormitory_i"}
	for _, stop := range stopList {
		request := httptest.NewRequest("GET", "/rest/shuttle/stop/"+stop, nil)
		app := test.InitApp()

		// Test
		var resp shuttle.StopItemResponse
		response, err := app.Test(request)
		if err != nil {
			t.Error(err)
		} else if response.StatusCode != 200 {
			t.Error("Status code is not 200")
		}
		if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
			t.Error(err)
		}
		if resp.Name == "" {
			t.Error("Stop name is empty")
		}
		if resp.Location.Latitude == 0 || resp.Location.Longitude == 0 {
			t.Error("Location is empty")
		}
		if resp.RouteList == nil {
			t.Error("Route list is nil")
		}
		for _, route := range resp.RouteList {
			if route.Name == "" {
				t.Error("Route name is empty")
			}
			if route.TimetableList == nil {
				t.Error("TimetableList is nil")
			}
			for _, timetable := range route.TimetableList {
				if timetable == "" {
					t.Error("Timetable is empty")
				}
			}
		}
	}
}

func TestShuttleRouteStopItem(t *testing.T) {
	// Request Object
	stopList := [6]string{"dormitory_o", "shuttlecock_o", "station", "terminal", "shuttlecock_i", "dormitory_i"}
	routeList := [3]string{"DH", "DY", "C"}
	for _, stop := range stopList {
		for _, route := range routeList {
			request := httptest.NewRequest("GET", "/rest/shuttle/stop/"+stop+"/route/"+route, nil)
			app := test.InitApp()

			// Test
			var resp shuttle.StopRouteItem
			response, err := app.Test(request)
			if err != nil {
				t.Error(err)
			} else if (stop == "station" && route == "DY") || (stop == "terminal" && route == "DH") {
				if response.StatusCode != 404 {
					t.Error("Status code is not 404")
				}
			} else if response.StatusCode != 200 {
				t.Error("Status code is not 200")
			} else {
				if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Error(err)
				}
				if resp.Name == "" {
					t.Error("Stop name is empty")
				}
				if resp.TimetableList == nil {
					t.Error("Timetable is nil")
				}
				for _, timetable := range resp.TimetableList {
					if timetable == "" {
						t.Error("Timetable is empty")
					}
				}
			}
		}
	}
}

func TestShuttleRouteStopTimetable(t *testing.T) {
	// Request Object
	stopList := [6]string{"dormitory_o", "shuttlecock_o", "station", "terminal", "shuttlecock_i", "dormitory_i"}
	routeList := [3]string{"DH", "DY", "C"}
	for _, stop := range stopList {
		for _, route := range routeList {
			request := httptest.NewRequest("GET", "/rest/shuttle/stop/"+stop+"/route/"+route+"/timetable", nil)
			app := test.InitApp()

			// Test
			var resp shuttle.StopRouteTimetableResponse
			response, err := app.Test(request)
			if err != nil {
				t.Error(err)
			} else if (stop == "station" && route == "DY") || (stop == "terminal" && route == "DH") {
				if response.StatusCode != 404 {
					t.Error("Status code is not 404")
				}
			} else if response.StatusCode != 200 {
				t.Error("Status code is not 200")
			} else {
				if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Error(err)
				}
				if resp.Name == "" {
					t.Error("Stop name is empty")
				}
				if resp.Weekdays == nil {
					t.Error("Timetable is nil")
				}
				for _, timetable := range resp.Weekends {
					if timetable == "" {
						t.Error("Timetable is empty")
					}
				}
				if resp.Weekends == nil {
					t.Error("Timetable is nil")
				}
				for _, timetable := range resp.Weekends {
					if timetable == "" {
						t.Error("Timetable is empty")
					}
				}
			}
		}
	}
}

func TestShuttleRouteStopArrival(t *testing.T) {
	// Request Object
	stopList := [6]string{"dormitory_o", "shuttlecock_o", "station", "terminal", "shuttlecock_i", "dormitory_i"}
	routeList := [3]string{"DH", "DY", "C"}
	for _, stop := range stopList {
		for _, route := range routeList {
			request := httptest.NewRequest("GET", "/rest/shuttle/stop/"+stop+"/route/"+route+"/arrival", nil)
			app := test.InitApp()

			// Test
			var resp shuttle.StopRouteArrivalResponse
			response, err := app.Test(request)
			if err != nil {
				t.Error(err)
			} else if (stop == "station" && route == "DY") || (stop == "terminal" && route == "DH") {
				if response.StatusCode != 404 {
					t.Error("Status code is not 404")
				}
			} else if response.StatusCode != 200 {
				t.Error("Status code is not 200")
			} else {
				if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Error(err)
				}
				if resp.Name == "" {
					t.Error("Stop name is empty")
				}
				if resp.ArrivalList == nil {
					t.Error("Arrival data is nil")
				}
				for _, arrival := range resp.ArrivalList {
					if reflect.TypeOf(arrival) != reflect.TypeOf(int64(0)) {
						t.Error("Arrival data is not int64")
					}
				}
			}
		}
	}
}
