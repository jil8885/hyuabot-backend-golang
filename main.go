package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/jil8885/hyuabot-backend-golang/common"
	"github.com/jil8885/hyuabot-backend-golang/kakao"
	"log"
	"time"
)

// 실제 서버 실행코드
func main()  {
	app := fiber.New()
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration: time.Minute,
		CacheControl: true,
	}))
	// 카카오 i 라우트
	kakaoUrl := app.Group("/kakao", kakao.Middleware)
	kakaoUrl.Post("/shuttle", kakao.Shuttle)
	kakaoUrl.Post("/shuttle/by-stop", kakao.ShuttleStop)
	kakaoUrl.Post("/subway", kakao.Subway)
	kakaoUrl.Post("/bus", kakao.Bus)
	kakaoUrl.Post("/food", kakao.Food)
	kakaoUrl.Post("/library", kakao.Library)

	// 휴아봇 앱 라우트
	// 공통 기능 라우트
	commonUrl := app.Group("/common", common.Middleware)
	commonUrl.Get("/library", common.Library)
	// Fatal Log 출력
	log.Fatal(app.Listen(":8080"))
}