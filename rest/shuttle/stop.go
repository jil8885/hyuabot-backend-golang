package shuttle

import "github.com/gofiber/fiber/v2"

// 셔틀버스 정류장 목록 조회
func GetShuttleStopList(c *fiber.Ctx) error {
	return c.SendString("GetShuttleStopList")
}

// 셔틀버스 정류장 항목 조회
func GetShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("GetShuttleStopItem")
}

// 셔틀버스 정류장 항목 추가
func PostShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("PostShuttleStopItem")
}

// 셔틀버스 정류장 항목 수정
func PutShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("PutShuttleStopItem")
}

// 셔틀버스 정류장 항목 삭제
func DeleteShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("DeleteShuttleStopItem")
}

// 셔틀버스 정류장 경유 노선 조회
func GetShuttleStopRoute(c *fiber.Ctx) error {
	return c.SendString("GetShuttleStopRoute")
}

// 셔틀버스 정류장별 시간표 조회
func GetShuttleStopRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("GetShuttleStopTimeTable")
}

// 셔틀버스 정류장별 도착 예정 시간 조회
func GetShuttleStopRouteArrivalTime(c *fiber.Ctx) error {
	return c.SendString("GetShuttleStopArrivalTime")
}
