package shuttle

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/shuttle"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/shuttle"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"time"
)

// 셔틀버스 시간표 일괄 조회
func GetShuttleTimeTable(c *fiber.Ctx) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var periodItem model.Period
	result := util.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&periodItem)

	var stopList []model.Stop
	result = util.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ?", periodItem.Type).
		Find(&stopList)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopTimetableListResponse(stopList))}

// 셔틀버스 도착 예정 시간 일괄 조회
func GetShuttleArrivalTime(c *fiber.Ctx) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var periodItem model.Period
	result := util.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&periodItem)

	var stopList []model.Stop
	result = util.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, now.Weekday() < 6).
		Find(&stopList)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopArrivalListResponse(stopList))
}
