package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jil8885/hyuabot-backend-golang/app"
	"github.com/jil8885/hyuabot-backend-golang/common"
	"github.com/jil8885/hyuabot-backend-golang/kakao"
	"log"
)

// 실제 서버 실행코드
func main() {
	server := fiber.New()
	server.Use(logger.New())

	// 카카오 i 라우트
	kakaoUrl := server.Group("/kakao", kakao.Middleware)
	kakaoUrl.Post("/shuttle", kakao.Shuttle)
	kakaoUrl.Post("/shuttle/all", kakao.GetAllShuttle)
	kakaoUrl.Post("/shuttle/by-stop", kakao.ShuttleStop)
	kakaoUrl.Post("/subway", kakao.Subway)
	kakaoUrl.Post("/bus", kakao.Bus)
	kakaoUrl.Post("/food", kakao.Food)
	kakaoUrl.Post("/library", kakao.Library)

	// 휴아봇 앱 라우트
	appUrl := server.Group("/app", app.Middleware)
	appUrl.Get("/shuttle", app.GetShuttleDeparture)
	appUrl.Post("/shuttle", app.GetShuttleDepartureByStopBackport) // Backport
	appUrl.Get("/shuttle/by-stop", app.GetShuttleStopInfoByStopWithParams)
	appUrl.Post("/shuttle/by-stop", app.GetShuttleStopInfoByStopBackport) // Backport
	appUrl.Get("/subway", app.GetSubwayDepartureWithParams)
	appUrl.Post("/subway", app.GetSubwayDepartureBackport) // Backport
	appUrl.Get("/bus", app.GetBusDeparture)
	appUrl.Get("/bus", app.GetBusDepartureByLineWithParams)
	appUrl.Post("/bus", app.GetBusDepartureByLineBackport) // Backport
	appUrl.Get("/bus/timetable", app.GetBusTimetableByRouteWithParams)
	appUrl.Post("/bus/timetable", app.GetBusTimetableByRouteBackport) // Backport
	appUrl.Get("/library", app.GetReadingRoomSeatByCampusWithParams)
	appUrl.Post("/library", app.GetReadingRoomSeatByCampusBackport) // Backport
	appUrl.Get("/food", app.GetFoodMenuByCampus)

	// 공통 기능 라우트
	commonUrl := server.Group("/common", common.Middleware)
	commonUrl.Get("/library", common.Library)
	commonUrl.Get("/food", common.Food)
	commonUrl.Get("/subway", common.Subway)

	// Fatal Log 출력
	log.Fatal(server.Listen(":8080"))
}
