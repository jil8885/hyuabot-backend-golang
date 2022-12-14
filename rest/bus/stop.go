package bus

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 버스 정류장 목록 조회
func GetBusStopList(c *fiber.Ctx) error {
	var busStopList []model.Stop
	var nameQuery = c.Query("name")
	if nameQuery == "" {
		util.DB.Database.Model(&model.Stop{}).
			Find(&busStopList)
	} else {
		util.DB.Database.Model(&model.Stop{}).
			Where("stop_name like ?", "%"+nameQuery+"%").
			Find(&busStopList)
	}
	return c.JSON(response.CreateStopListResponse(busStopList))
}

// 버스 정류장 항목 조회
func GetBusStopItem(c *fiber.Ctx) error {
	return c.SendString("GetBusStopItem")
}
