package cafeteria

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	cafeteria := router.Group("/cafeteria")
	cafeteria.Get("/campus/:campus_id/restaurant", GetRestaurantList)
	cafeteria.Get("/campus/:campus_id/restaurant/:restaurant_id", GetRestaurantItem)
	cafeteria.Get("/campus/:campus_id/menu", GetCafeteriaMenu)
}
