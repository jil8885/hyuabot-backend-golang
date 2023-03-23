package shuttle

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetShuttleStopList(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetShuttleStopList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/shuttle/stop", GetShuttleStopList)
	request := httptest.NewRequest("GET", "/rest/shuttle/stop", nil)
	res, err := app.Test(request)
	test.Nil(err)
	test.Equal(200, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	test.Nil(err)
	var obj shuttle.StopListResponse
	err = json.Unmarshal(body, &obj)
	test.Nil(err)
	test.IsType([]shuttle.StopListItem{}, obj.Stop)
	test.Greater(len(obj.Stop), 0, "There should be at least one shuttle route")
	for _, stop := range obj.Stop {
		test.IsType(shuttle.StopListItem{}, stop)
		test.IsType("", stop.Name)
		test.IsType(shuttle.StopLocation{}, stop.Location)
		test.IsType(0.0, stop.Location.Latitude)
		test.IsType(0.0, stop.Location.Longitude)
	}
}
