package shuttle

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"

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
	lunarYear, lunarMonth, lunarDay := carbon.Now().Lunar().Date()

	var holidayItem model.Holiday
	utils.DB.Database.Model(&model.Holiday{}).
		Where("(holiday_date = ? and calendar_type = ?) or (holiday_date = ? and calendar_type = ?)",
			now.Format("2006-01-02"), "solar", fmt.Sprintf("%d-%d-%d", lunarYear, lunarMonth, lunarDay), "lunar").
		First(&holidayItem)
	var periodItem model.Period
	result := utils.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&periodItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var holiday = now.Weekday() != time.Saturday && now.Weekday() != time.Sunday
	if holidayItem.HolidayType == "weekends" {
		holiday = true
	}

	var stopList []model.Stop
	result = utils.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, now, holiday).
		Preload("RouteList.ShuttleRoute").
		Find(&stopList)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopArrivalListResponse(holidayItem.HolidayType, stopList))
}
