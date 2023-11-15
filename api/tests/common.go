package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
	"github.com/hyuabot-developers/hyuabot-backend-golang/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http/httptest"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/hyuabot-developers/hyuabot-backend-golang/api/route"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
)

func setup() *fiber.App {
	app := fiber.New()
	route.SetupRouterV1(app)
	return app
}

func setupDatabase() {
	database.ConnectDB(os.Getenv("DB_NAME"))
	database.ConnectRedis()
}

func tearDownDatabase() {
	// Remove all bus data
	_, _ = database.DB.Exec("DELETE FROM bus_realtime")
	_, _ = database.DB.Exec("DELETE FROM bus_timetable")
	_, _ = database.DB.Exec("DELETE FROM bus_route_stop")
	_, _ = database.DB.Exec("DELETE FROM bus_route")
	_, _ = database.DB.Exec("DELETE FROM bus_stop")
	// Remove all commute shuttle data
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_timetable")
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_route")
	_, _ = database.DB.Exec("DELETE FROM commute_shuttle_stop")
	// Remove all shuttle data
	_, _ = database.DB.Exec("DELETE FROM shuttle_timetable")
	_, _ = database.DB.Exec("DELETE FROM shuttle_route_stop")
	_, _ = database.DB.Exec("DELETE FROM shuttle_route")
	_, _ = database.DB.Exec("DELETE FROM shuttle_stop")
	_, _ = database.DB.Exec("DELETE FROM shuttle_holiday")
	_, _ = database.DB.Exec("DELETE FROM shuttle_period")
	_, _ = database.DB.Exec("DELETE FROM shuttle_period_type")
	// Remove all subway data
	_, _ = database.DB.Exec("DELETE FROM subway_realtime")
	_, _ = database.DB.Exec("DELETE FROM subway_timetable")
	_, _ = database.DB.Exec("DELETE FROM subway_route_station")
	_, _ = database.DB.Exec("DELETE FROM subway_route")
	_, _ = database.DB.Exec("DELETE FROM subway_station")
	// Remove all meal data
	_, _ = database.DB.Exec("DELETE FROM menu")
	_, _ = database.DB.Exec("DELETE FROM restaurant")
	// Remove all reading room data
	_, _ = database.DB.Exec("DELETE FROM reading_room")
	// Remove all campus data
	_, _ = database.DB.Exec("DELETE FROM campus")
	// Remove all notice data
	_, _ = database.DB.Exec("DELETE FROM notices")
	_, _ = database.DB.Exec("DELETE FROM notice_category")
	// Remove all users
	_, _ = database.DB.Exec("DELETE FROM admin_user")
}

func createAdminUser() {
	// Insert test user
	hashedPassword, _ := utils.HashPassword("test")
	user := models.AdminUser{
		UserID:   "test",
		Password: hashedPassword,
		Name:     "test",
		Email:    "test@email.com",
		Phone:    "010-1234-5678",
		Active:   true,
	}
	ctx := context.Background()
	_ = user.Insert(ctx, database.DB, boil.Infer())
}

func loginWithAdminUser(app *fiber.App) string {
	admin := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "test",
		Password: "test",
	}

	body, _ := json.Marshal(admin)
	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)

	var result responses.TokenResponse
	_ = json.NewDecoder(response.Body).Decode(&result)
	return result.AccessToken
}
