package shuttle

import (
	"time"

	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 셔틀버스 노선 목록 조회
func GetShuttleRouteList(c *fiber.Ctx) error {
	var routeList []model.RouteItem
	utils.DB.Database.Model(&model.Route{}).Find(&routeList)
	return c.JSON(response.CreateRouteListResponse(routeList))
}

// 셔틀버스 노선 항목 조회
func GetShuttleRouteItem(c *fiber.Ctx) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var periodItem model.Period
	result := utils.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&periodItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var routeItem model.Route
	result = utils.DB.Database.Model(&model.Route{}).
		Preload("StopList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Where("shuttle_route.route_name = ?", c.Params("route_id")).
		First(&routeItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateRouteItemResponse(routeItem))
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
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var periodItem model.Period
	result := utils.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&periodItem)

	var routeItem model.Route
	result = utils.DB.Database.Model(&model.Route{}).
		Preload("StopList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Where("shuttle_route.route_name = ?", c.Params("route_id")).
		First(&routeItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateRouteLocationResponse(routeItem))
}
