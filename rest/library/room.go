package library

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/library"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/library"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 캠퍼스별 열람실 목록 조회
func GetLibraryRoomList(c *fiber.Ctx) error {
	var roomList []model.Room
	util.DB.Database.Model(&model.Room{}).
		Where("campus_id = ?", c.Params("campus_id")).
		Find(&roomList)
	return c.JSON(response.CreateRoomListResponse(roomList))
}

// 열람실 항목 조회
func GetLibraryRoomItem(c *fiber.Ctx) error {
	return c.SendString("GetLibraryRoomItem")
}
