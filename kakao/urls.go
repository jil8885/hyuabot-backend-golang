package kakao

import "github.com/gofiber/fiber/v2"

// 카카오 i 용 url handler
func Middleware(c *fiber.Ctx) error {
	c.Accepts("application/json") // json 형식으로만 요청 가능
	return c.Next()
}

// 카카오 i 셔틀 도착 정보 제공
func Shuttle(c *fiber.Ctx) error {
	return c.SendString("카카오 i 셔틀 도착 정보")
}

// 카카오 i 셔틀 정류장 정보 제공
func ShuttleStop(c *fiber.Ctx) error {
	return c.SendString("카카오 i 셔틀 정류장 정보")
}

// 카카오 i 전철 도착 정보 제공
func Subway(c *fiber.Ctx) error {
	return c.SendString("카카오 i 전철 도착 정보")
}

// 카카오 i 버스 도착 정보 제공
func Bus(c *fiber.Ctx) error {
	return c.SendString(parseAnswer(c))
}

// 카카오 i 학식 정보 제공
func Food(c *fiber.Ctx) error {
	return c.SendString("카카오 i 학식 정보")
}

// 카카오 i 열람실 정보 제공
func Library(c *fiber.Ctx) error {
	return c.SendString("카카오 i 열람실 정보")
}

// 카카오톡을 통해 넘어온 데이터 중 사용자의 발화 Parse
func parseAnswer(c *fiber.Ctx) string {
	model := new(UserMessage)
	if err := c.BodyParser(model); err != nil{
		return err.Error()
	}
	return model.Request.Message
}