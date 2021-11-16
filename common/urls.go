package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jil8885/hyuabot-backend-golang/food"
	"github.com/jil8885/hyuabot-backend-golang/library"
	"github.com/jil8885/hyuabot-backend-golang/subway"
)

func Middleware(c *fiber.Ctx) error {
	return c.Next()
}

func Library(c *fiber.Ctx) error {
	library.FetchLibrary()
	return c.SendString("도서관 크롤링")
}

func Food(c *fiber.Ctx) error {
	food.FetchFoodMenu()
	return c.SendString("학식 크롤링")
}

func Subway(c *fiber.Ctx) error{
	subway.FetchSubwayRealtime(0, 1004)
	return c.SendString("지하철 크롤링")
}

func Bus(c *fiber.Ctx) error{
	return c.SendString("버스 크롤링")
}

