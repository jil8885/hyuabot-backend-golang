package library

import "github.com/gofiber/fiber/v2"

// 캠퍼스별 열람실 목록 조회
func GetLibraryRoomList(c *fiber.Ctx) error {
	return c.SendString("GetLibraryRoomList")
}

// 열람실 항목 조회
func GetLibraryRoomItem(c *fiber.Ctx) error {
	return c.SendString("GetLibraryRoomItem")
}
