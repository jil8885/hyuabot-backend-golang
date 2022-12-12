package cafeteria

import "github.com/gofiber/fiber/v2"

// 캠퍼스별 학식 식당 목록 조회
func GetRestaurantList(c *fiber.Ctx) error {
	return c.SendString("GetRestaurantList")
}

// 식당 항목 조회
func GetRestaurantItem(c *fiber.Ctx) error {
	return c.SendString("GetRestaurantItem")
}
