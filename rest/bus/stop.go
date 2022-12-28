package bus

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/bus"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/bus"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 평일/주말 구분 함수
func isWeekend(now time.Time) string {
	if now.Weekday() == time.Saturday {
		return "saturday"
	} else if now.Weekday() == time.Sunday {
		return "sunday"
	}
	return "weekdays"
}

// 버스 정류장 목록 조회
func GetBusStopList(c *fiber.Ctx) error {
	var busStopList []model.Stop
	var nameQuery = c.Query("name")
	if nameQuery == "" {
		utils.DB.Database.Model(&model.Stop{}).
			Find(&busStopList)
	} else {
		utils.DB.Database.Model(&model.Stop{}).
			Where("stop_name like ?", "%"+nameQuery+"%").
			Find(&busStopList)
	}
	return c.JSON(response.CreateStopListResponse(busStopList))
}

// 버스 정류장 항목 조회
func GetBusStopItem(c *fiber.Ctx) error {
	// 기준 날짜 로딩
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var busStopItem model.Stop
	result := utils.DB.Database.Model(&model.Stop{}).
		Preload("RouteList.RouteItem").
		Preload("RouteList.StartStop").
		Preload("RouteList.TimetableList", "weekday = ? and (departure_time > ? or departure_time < '04:00:00')",
			isWeekend(now),
			fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())).
		Preload("RouteList.RealtimeList").
		Where("stop_id = ?", c.Params("stop_id")).
		First(&busStopItem)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(response.CreateStopItemResponse(busStopItem))
}
