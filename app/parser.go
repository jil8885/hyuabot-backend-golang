package app

import "github.com/gofiber/fiber/v2"

func parseShuttleStop(c *fiber.Ctx) string {
	model := new(ShuttleStopRequest)
	if err := c.BodyParser(model); err != nil{
		return err.Error()
	}
	return model.BusStop
}

func parseCampus(c *fiber.Ctx) string {
	model := new(CampusRequest)
	if err := c.BodyParser(model); err != nil{
		return err.Error()
	}
	return model.Campus
}

func parseBusRouteID(c *fiber.Ctx) string {
	model := new(BusRouteRequest)
	if err := c.BodyParser(model); err != nil{
		return err.Error()
	}
	return model.Route
}
