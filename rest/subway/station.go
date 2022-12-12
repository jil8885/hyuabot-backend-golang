package subway

import "github.com/gofiber/fiber/v2"

// 전철역 목록 조회
func GetStationList(c *fiber.Ctx) error {
	return c.SendString("GetStationList")
}

// 전철역 항목 조회
func GetStationItem(c *fiber.Ctx) error {
	return c.SendString("GetStationItem")
}

// 전철역 항목 추가
func PostStationItem(c *fiber.Ctx) error {
	return c.SendString("PostStationItem")
}

// 전철역 항목 수정
func PutStationItem(c *fiber.Ctx) error {
	return c.SendString("PutStationItem")
}

// 전철역 항목 삭제
func DeleteStationItem(c *fiber.Ctx) error {
	return c.SendString("DeleteStationItem")
}

// 전철역 도착 정보 조회
func GetStationArrival(c *fiber.Ctx) error {
	return c.SendString("GetStationArrival")
}

// 전철역 시간표 일괄 조회
func GetStationTimeTable(c *fiber.Ctx) error {
	return c.SendString("GetStationTimeTable")
}
