package library

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	cafeteria := router.Group("/library")
	cafeteria.Get("/campus/:campus_id", GetLibraryRoomList)
	cafeteria.Get("/campus/:campus_id/room/:room_id", GetLibraryRoomItem)
}
