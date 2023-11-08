package utils

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := CheckTokenIsValid(c.Request())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
		}
		return c.Next()
	}
}
