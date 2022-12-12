package shuttle

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 셔틀버스 노선 목록 조회
func GetShuttleRouteList(c *fiber.Ctx) error {
	var routeList []model.RouteItem
	util.DB.Database.Model(&model.Route{}).Find(&routeList)
	return c.JSON(response.CreateRouteListResponse(routeList))
}

// 셔틀버스 노선 항목 조회
func GetShuttleRouteItem(c *fiber.Ctx) error {
	return c.SendString("GetShuttleRouteItem")
}

// 셔틀버스 노선 항목 추가
func PostShuttleRouteItem(c *fiber.Ctx) error {
	return c.SendString("PostShuttleRouteItem")
}

// 셔틀버스 노선 항목 수정
func PutShuttleRouteItem(c *fiber.Ctx) error {
	return c.SendString("PutShuttleRouteItem")
}

// 셔틀버스 노선 항목 삭제
func DeleteShuttleRouteItem(c *fiber.Ctx) error {
	return c.SendString("DeleteShuttleRouteItem")
}

// 셔틀버스 노선별 위치 조회
func GetShuttleRouteLocation(c *fiber.Ctx) error {
	return c.SendString("GetShuttleRouteLocation")
}
