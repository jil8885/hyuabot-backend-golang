package shuttle

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/test"
)

func TestShuttleRouteList(t *testing.T) {
	// Request Object
	request := httptest.NewRequest("GET", "/rest/shuttle/route", nil)
	app := test.InitApp()

	// Test
	var resp shuttle.RouteListResponse
	response, err := app.Test(request)
	if err != nil {
		t.Error(err)
	} else if response.StatusCode != 200 {
		t.Error("Status code is not 200")
	}
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	if resp.Route == nil {
		t.Error("Response is nil")
	}
	for _, route := range resp.Route {
		if route.Name == "" {
			t.Error("RouteID is empty")
		}
		if route.Description.Korean == "" || route.Description.English == "" {
			t.Error("Route Description is empty")
		}
	}
}

func TestShuttleRouteItem(t *testing.T) {
	// Request Object
	routeList := [3]string{"DH", "DY", "C"}
	for _, route := range routeList {
		request := httptest.NewRequest("GET", "/rest/shuttle/route/"+route, nil)
		app := test.InitApp()

		// Test
		var resp shuttle.RouteItemResponse
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
			t.Error("RouteID is empty")
		}
		if resp.Description.Korean == "" || resp.Description.English == "" {
			t.Error("Route Description is empty")
		}
		if resp.StopList == nil {
			t.Error("StopList is nil")
		}
		for _, stop := range resp.StopList {
			if stop.Name == "" {
				t.Error("StopName is empty")
			}
			if stop.TimetableList == nil {
				t.Error("TimetableList is nil")
			}
			for _, timetable := range stop.TimetableList {
				if timetable == "" {
					t.Error("Timetable is empty")
				}
			}
		}
	}
}
