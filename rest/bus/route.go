package bus

import (
	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 버스 노선 목록 조회
func GetBusRouteList(c *fiber.Ctx) error {
	var busRouteList []model.Route
	var nameQuery = c.Query("name")
	if nameQuery == "" {
		utils.DB.Database.Model(&model.Route{}).
			Find(&busRouteList)
	} else {
		utils.DB.Database.Model(&model.Route{}).
			Where("route_name like ?", "%"+nameQuery+"%").
			Find(&busRouteList)
	}
	return c.JSON(response.CreateRouteListResponse(busRouteList))
}

// 버스 노선 항목 조회
func GetBusRouteItem(c *fiber.Ctx) error {
	var busRouteItem model.Route
	var routeID = c.Params("route_id")
	result := utils.DB.Database.Model(&model.Route{}).
		Preload("StartStop").
		Preload("EndStop").
		Where("route_id = ?", routeID).
		First(&busRouteItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateRouteItemResponse(busRouteItem))
}

// 버스 노선 항목 추가
func PostBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("PostBusRouteItem")
}

// 버스 노선 항목 수정
func PutBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("PutBusRouteItem")
}

// 버스 노선 항목 삭제
func DeleteBusRouteItem(c *fiber.Ctx) error {
	return c.SendString("DeleteBusRouteItem")
}

// 버스 노선별 시간표 조회
func GetBusRouteTimeTable(c *fiber.Ctx) error {
	var busRouteStopItem model.RouteStop
	routeID := c.Params("route_id")
	startStopID := c.Params("stop_id")
	result := utils.DB.Database.Model(&model.RouteStop{}).
		Preload("TimetableList").
		Where("route_id = ? and start_stop_id = ?", routeID, startStopID).
		First(&busRouteStopItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateRouteTimetableResponse(busRouteStopItem.TimetableList))
}

// 버스 노선별 시간표 삭제
func DeleteBusRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("DeleteBusRouteTimeTable")
}

// 버스 노선별 시간표 추가
func PostBusRouteTimeTable(c *fiber.Ctx) error {
	return c.SendString("PostBusRouteTimeTable")
}
