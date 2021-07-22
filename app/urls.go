package app

import "github.com/gofiber/fiber/v2"

// App 용 url handler
func Middleware(c *fiber.Ctx) error {
	c.Accepts("application/json") // json 형식으로만 요청 가능
	return c.Next()
}

func GetShuttleDeparture(c *fiber.Ctx) error {
	return c.SendString("전체 정류장 셔틀 도착 정보 조회")
}

func GetShuttleDepartureByStop(c *fiber.Ctx) error {
	return c.SendString("정류장 별 셔틀 도착 정보 조회")
}

func GetShuttleStopInfoByStop(c *fiber.Ctx) error {
	return c.SendString("정류장 별 정보 조회")
}

func GetSubwayDeparture(c *fiber.Ctx) error {
	return c.SendString("전철 도착 정보 조회")
}

func GetBusDeparture(c *fiber.Ctx) error {
	return c.SendString("전체 버스 도착 정보")
}

func GetBusDepartureByLine(c *fiber.Ctx) error {
	return c.SendString("노선별 버스 도착 정보")
}

func GetBusTimetableByRoute(c *fiber.Ctx) error  {
	return c.SendString("노선별 버스 시간표")
}

func GetReadingRoomSeatByCampus(c *fiber.Ctx) error{
	return c.SendString("캠퍼스별 버스 시간표")
}

func PushNotificationByRoom(c *fiber.Ctx) error  {
	return c.SendString("열람실 별 잔여좌석 알림")
}

func GetFoodMenuByCampus(c *fiber.Ctx) error  {
	return c.SendString("식당별 메뉴")
}
