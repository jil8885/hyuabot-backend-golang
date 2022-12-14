package cafeteria

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	model "github.com/hyuabot-developers/hyuabot-backend-golang/model/cafeteria"
	response "github.com/hyuabot-developers/hyuabot-backend-golang/response/cafeteria"
	"github.com/hyuabot-developers/hyuabot-backend-golang/util"
	"time"
)

// 캠퍼스별 학식 식당 목록 조회
func GetRestaurantList(c *fiber.Ctx) error {
	// 기준 날짜 로딩
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	date := c.Query("date")
	if date == "" {
		date = fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	}
	// 조식, 중식, 석식 시간 로딩
	timeType := c.Query("time")
	if timeType == "" {
		switch {
		case now.Hour() < 10:
			timeType = "조식"
		case now.Hour() > 15:
			timeType = "석식"
		default:
			timeType = "중식"
		}
	}

	var restaurantList []model.RestaurantItem
	util.DB.Database.Model(&model.Restaurant{}).
		Preload("MenuList", "feed_date = ? and time_type like ?",
			fmt.Sprintf(date),
			fmt.Sprintf("%%%s%%", timeType)).
		Where("campus_id = ?", c.Params("campus_id")).
		Find(&restaurantList)
	return c.JSON(response.CreateRestaurantListResponse(restaurantList))
}

// 식당 항목 조회
func GetRestaurantItem(c *fiber.Ctx) error {
	// 기준 날짜 로딩
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	date := c.Query("date")
	if date == "" {
		date = fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	}
	// 조식, 중식, 석식 시간 로딩
	timeType := c.Query("time")
	if timeType == "" {
		switch {
		case now.Hour() < 10:
			timeType = "조식"
		case now.Hour() > 15:
			timeType = "석식"
		default:
			timeType = "중식"
		}
	}

	var restaurantItem model.RestaurantItem
	util.DB.Database.Model(&model.Restaurant{}).
		Preload("MenuList", "feed_date = ? and time_type like ?",
			fmt.Sprintf(date),
			fmt.Sprintf("%%%s%%", timeType)).
		Where("campus_id = ? and restaurant_id = ?", c.Params("campus_id"), c.Params("restaurant_id")).
		Find(&restaurantItem)
	return c.JSON(response.CreateRestaurantItemResponse(restaurantItem))
}
