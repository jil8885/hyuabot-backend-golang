package library

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	library := router.Group("/library")
	library.Get("/campus/:campus_id", GetLibraryRoomList)
	library.Get("/campus/:campus_id/room/:room_id", GetLibraryRoomItem)
}
