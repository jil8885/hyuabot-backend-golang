package bus

import "github.com/gofiber/fiber/v2"

// 버스 노선 목록 조회
func GetBusRouteList(c *fiber.Ctx) error {
	return c.SendString("GetBusRouteList")
}

// 버스 노선 항목 조회
func GetBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("GetBusRouteItem")
}

// 버스 노선 항목 추가
func PostBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("PostBusRouteItem")
}

// 버스 노선 항목 수정
func PutBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("PutBusRouteItem")
}

// 버스 노선 항목 삭제
func DeleteBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("DeleteBusRouteItem")
}

// 버스 노선별 시간표 조회
func GetBusRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("GetBusRouteTimeTable")
}

// 버스 노선별 시간표 삭제
func DeleteBusRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("DeleteBusRouteTimeTable")
}

// 버스 노선별 시간표 추가
func PostBusRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("PostBusRouteTimeTable")
}
