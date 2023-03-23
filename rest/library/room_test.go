package library

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hyuabot-developers/hyuabot-backend-golang/response/library"

	"github.com/gofiber/fiber/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"github.com/stretchr/testify/assert"
)

func TestGetLibraryRoomList(t *testing.T) {
	test := assert.New(t)
	t.Log("TestGetLibraryRoomList")
	util.ConnectDB()
	app := fiber.New()
	app.Get("/rest/library/campus/:campus_id", GetLibraryRoomList)
	for i := 1; i <= 2; i++ {
		request := httptest.NewRequest("GET", fmt.Sprintf("/rest/library/campus/%d", i), nil)
		res, err := app.Test(request)
		test.Nil(err)
		body, err := io.ReadAll(res.Body)
		test.Nil(err)
		test.Equal(200, res.StatusCode)
		var obj library.RoomListResponse
		err = json.Unmarshal(body, &obj)
		test.Nil(err)
		test.IsType([]library.RoomItemResponse{}, obj.RoomList)
		test.Greater(len(obj.RoomList), 0, "There should be at least one library room")
		for _, room := range obj.RoomList {
			test.IsType(library.RoomItemResponse{}, room)
			test.IsType(0, room.RoomID)
			test.IsType("", room.Name)
			test.IsType(0, room.Total)
			test.IsType(0, room.Active)
			test.IsType(0, room.Occupied)
			test.IsType(0, room.Available)
			test.IsType("", room.LastUpdate)
		}
	}
}
