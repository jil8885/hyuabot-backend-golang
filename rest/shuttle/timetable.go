package shuttle

import (
	"time"

	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 셔틀버스 시간표 일괄 조회
func GetShuttleTimeTable(c *fiber.Ctx) error {
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
	var stopList []model.Stop
	result = utils.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ?", periodItem.Type).
		Preload("RouteList.ShuttleRoute").
		Find(&stopList)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopTimetableListResponse(stopList))
}

// 셔틀버스 도착 예정 시간 일괄 조회
func GetShuttleArrivalTime(c *fiber.Ctx) error {
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

	var stopList []model.Stop
	result = utils.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() != time.Saturday && now.Weekday() != time.Sunday).
		Preload("RouteList.ShuttleRoute").
		Find(&stopList)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopArrivalListResponse(stopList))
}
