package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/jil8885/hyuabot-backend-golang/app"
	"github.com/jil8885/hyuabot-backend-golang/common"
	"github.com/jil8885/hyuabot-backend-golang/kakao"
	"log"
	"time"
)

// 실제 서버 실행코드
func main()  {
	server := fiber.New()
	server.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration: time.Minute,
		CacheControl: true,
	}))
	// 카카오 i 라우트
	kakaoUrl := server.Group("/kakao", kakao.Middleware)
	kakaoUrl.Post("/shuttle", kakao.Shuttle)
	kakaoUrl.Post("/shuttle/by-stop", kakao.ShuttleStop)
	kakaoUrl.Post("/subway", kakao.Subway)
	kakaoUrl.Post("/bus", kakao.Bus)
	kakaoUrl.Post("/food", kakao.Food)
	kakaoUrl.Post("/library", kakao.Library)

	// 휴아봇 앱 라우트
	appUrl := server.Group("/app", app.Middleware)
	appUrl.Get("/shuttle", app.GetShuttleDeparture)
	appUrl.Post("/shuttle", app.GetShuttleDepartureByStop)
	appUrl.Post("/shuttle/by-stop", app.GetShuttleStopInfoByStop)
	appUrl.Post("/subway", app.GetSubwayDeparture)
	appUrl.Get("/bus", app.GetBusDeparture)
	appUrl.Post("/bus", app.GetBusDepartureByLine)
	appUrl.Post("/bus/timetable", app.GetBusTimetableByRoute)
	appUrl.Post("/library", app.GetReadingRoomSeatByCampus)
	appUrl.Post("/food", app.GetFoodMenuByCampus)

	// 공통 기능 라우트
	commonUrl := server.Group("/common", common.Middleware)
	commonUrl.Get("/library", common.Library)
	commonUrl.Get("/food", common.Food)

	// Fatal Log 출력
	log.Fatal(server.Listen(":8080"))
}