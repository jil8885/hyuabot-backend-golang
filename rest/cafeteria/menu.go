package cafeteria

import "github.com/gofiber/fiber/v2"

// 현재 학식 메뉴 조회
func GetCafeteriaMenu(c *fiber.Ctx) error {
	return c.SendString("GetCafeteriaMenu")
}
