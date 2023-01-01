package shuttle

import (
	"fmt"

	"github.com/golang-module/carbon/v2"

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

func GetShuttleRouteTimetable(c *fiber.Ctx, dataType string) error {
	dateQuery := c.Query("date")
	now := carbon.Now("Asia/Seoul")
	date := now
	if dateQuery != "" {
		date = carbon.Parse(dateQuery, "Asia/Seoul").SetTime(now.Hour(), now.Minute(), now.Second())
	}
	lunarYear, lunarMonth, lunarDay := date.Lunar().Date()

	var holidayItem model.Holiday
	utils.DB.Database.Model(&model.Holiday{}).
		Where("(holiday_date = ? and calendar_type = ?) or (holiday_date = ? and calendar_type = ?)",
			date.ToDateString(), "solar", fmt.Sprintf("%d-%d-%d", lunarYear, lunarMonth, lunarDay), "lunar").
		First(&holidayItem)

	var periodItem model.Period
	result := utils.DB.Database.Model(&model.Period{}).
		Where("period_start <= ?", date).
		Where("period_end >= ?", date).
		First(&periodItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var holiday = date.IsWeekday()
	if holidayItem.HolidayType == "weekends" {
		holiday = true
	}
	var routeItem model.Route
	result = utils.DB.Database.Model(&model.Route{}).
		Preload("StopList.TimetableList", "period_type = ? and departure_time >= ? and weekday = ?",
			periodItem.Type, date, holiday).
		Where("shuttle_route.route_name = ?", c.Params("route_id")).
		First(&routeItem)
	// 해당 노선 ID가 존재하지 않는 경우
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if dataType == "arrival" {
		return c.JSON(response.CreateRouteItemResponse(holidayItem.HolidayType, holiday, date, routeItem))
	}
	return c.JSON(response.CreateRouteLocationResponse(holidayItem.HolidayType, routeItem))
}

// 셔틀버스 노선 항목 조회
func GetShuttleRouteItem(c *fiber.Ctx) error {
	return GetShuttleRouteTimetable(c, "arrival")
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
	return GetShuttleRouteTimetable(c, "location")
}
