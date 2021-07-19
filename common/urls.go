package common

import "github.com/gofiber/fiber/v2"

func Middleware(c *fiber.Ctx) error {
	return c.Next()
}

func Library(c *fiber.Ctx) error {
	return c.SendString("도서관 크롤링")
}