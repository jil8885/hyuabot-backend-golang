package subway

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/subway"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/subway"
	utils "github.com/hyuabot-developers/hyuabot-backend-golang/util"
)

// 평일/주말 구분 함수
func isWeekend(now time.Time) string {
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return "weekends"
	}
	return "weekdays"
}

// 전철역 목록 조회
func GetStationList(c *fiber.Ctx) error {
	var stopList []model.RouteStationListItem
	nameQuery := c.Query("name")
	if nameQuery == "" {
		utils.DB.Database.Model(&model.RouteStation{}).Find(&stopList)
	} else {
		utils.DB.Database.
			Model(&model.RouteStation{}).
			Where("station_name like ?", "%"+nameQuery+"%").
			Find(&stopList)
	}
	return c.JSON(response.CreateStationListResponse(stopList))
}

// 전철역 항목 조회
func GetStationItem(c *fiber.Ctx) error {
	// 기준 날짜 로딩
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var stationItem model.RouteStationItem
	stationID := c.Params("station_id")
	utils.DB.Database.
		Model(&model.RouteStation{}).
		Preload("RealtimeList.TerminalStation").
		Preload("TimetableList", "weekday = ? and departure_time > ?",
			isWeekend(now),
			fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())).
		Preload("TimetableList.TerminalStation").
		Where("station_id = ?", stationID).
		First(&stationItem)
	return c.JSON(response.CreateStationItemResponse(stationItem))
}

// 전철역 항목 추가
func PostStationItem(c *fiber.Ctx) error {
	return c.SendString("PostStationItem")
}

// 전철역 항목 수정
func PutStationItem(c *fiber.Ctx) error {
	return c.SendString("PutStationItem")
}

// 전철역 항목 삭제
func DeleteStationItem(c *fiber.Ctx) error {
	return c.SendString("DeleteStationItem")
}

// 전철역 도착 정보 조회
func GetStationArrival(c *fiber.Ctx) error {
	// 기준 날짜 로딩
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	var stationItem model.RouteStationItem
	stationID := c.Params("station_id")
	utils.DB.Database.
		Model(&model.RouteStation{}).
		Preload("RealtimeList.TerminalStation").
		Preload("TimetableList", "weekday = ? and departure_time > ?",
			isWeekend(now),
			fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())).
		Preload("TimetableList.TerminalStation").
		Where("station_id = ?", stationID).
		First(&stationItem)
	return c.JSON(response.CreateArrivalItemResponse(stationItem))
}

// 전철역 시간표 일괄 조회
func GetStationTimeTable(c *fiber.Ctx) error {
	var stationItem model.RouteStationItem
	stationID := c.Params("station_id")
	utils.DB.Database.
		Model(&model.RouteStation{}).
		Preload("RealtimeList.TerminalStation").
		Preload("TimetableList.TerminalStation").
		Where("station_id = ?", stationID).
		First(&stationItem)
	return c.JSON(response.CreateStationTimetableResponse(stationItem.TimetableList))
}
