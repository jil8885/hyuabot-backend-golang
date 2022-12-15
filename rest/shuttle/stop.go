package shuttle

import (
	"time"

	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 셔틀버스 정류장 목록 조회
func GetShuttleStopList(c *fiber.Ctx) error {
	var stopList []model.StopItem
	utils.DB.Database.Model(&model.Stop{}).Find(&stopList)
	return c.JSON(response.CreateStopListResponse(stopList))
}

// 셔틀버스 정류장 항목 조회
func GetShuttleStopItem(c *fiber.Ctx) error {
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

	var stopItem model.Stop
	result = utils.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Where("shuttle_stop.stop_name = ?", c.Params("stop_id")).
		First(&stopItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopItemResponse(stopItem))
}

// 셔틀버스 정류장 항목 추가
func PostShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("PostShuttleStopItem")
}

// 셔틀버스 정류장 항목 수정
func PutShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("PutShuttleStopItem")
}

// 셔틀버스 정류장 항목 삭제
func DeleteShuttleStopItem(c *fiber.Ctx) error {
	return c.SendString("DeleteShuttleStopItem")
}

// 셔틀버스 정류장 경유 노선 조회
func GetShuttleStopRoute(c *fiber.Ctx) error {
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

	var stopRouteItem model.RouteStop
	result = utils.DB.Database.Model(&model.RouteStop{}).
		Preload("TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Where("stop_name = ? and route_name = ?",
			c.Params("stop_id"), c.Params("route_id")).
		First(&stopRouteItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopRouteItem(stopRouteItem))
}

// 셔틀버스 정류장별 시간표 조회
func GetShuttleStopRouteTimeTable(c *fiber.Ctx) error {
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

	var stopRouteItem model.RouteStop
	result = utils.DB.Database.Model(&model.RouteStop{}).
		Preload("TimetableList", "period_type = ?", periodItem.Type).
		Where("stop_name = ? and route_name = ?", c.Params("stop_id"), c.Params("route_id")).
		First(&stopRouteItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopRouteTimetableResponse(stopRouteItem))
}

// 셔틀버스 정류장별 도착 예정 시간 조회
func GetShuttleStopRouteArrivalTime(c *fiber.Ctx) error {
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

	var stopRouteItem model.RouteStop
	result = utils.DB.Database.Model(&model.RouteStop{}).
		Preload("TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Where("stop_name = ? and route_name = ?",
			c.Params("stop_id"), c.Params("route_id")).
		First(&stopRouteItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopRouteArrivalItem(stopRouteItem))
}
