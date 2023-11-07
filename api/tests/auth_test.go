package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
	"github.com/hyuabot-developers/hyuabot-backend-golang/utils"
)

func TestSignUp(t *testing.T) {
	setupDatabase()
	test := assert.New(t)
	testCases := []struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}{
		// Provide invalid JSON
		{
			UserName: "test",
			Password: "test",
			Name:     "test",
			Email:    "test@email.com",
		},
		// Provide successful request
		{
			UserName: "test",
			Password: "test",
			Name:     "test",
			Email:    "test@email.com",
			Phone:    "010-1234-5678",
		},
		// Provide duplicated username
		{
			UserName: "test",
			Password: "test",
			Name:     "test",
			Email:    "test@email.com",
			Phone:    "010-1234-5678",
		},
	}

	expectedStatusCodes := []int{422, 201, 409}

	for index, testCase := range testCases {
		app := setup()
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)
	}
	tearDownDatabase()
}

func TestLogin(t *testing.T) {
	setupDatabase()
	// Insert test user
	hashedPassword, _ := utils.HashPassword("test")
	user := models.AdminUser{
		UserID:   "test",
		Password: hashedPassword,
		Name:     "test",
		Email:    "test@email.com",
		Phone:    "010-1234-5678",
		Active:   false,
	}
	ctx := context.Background()
	_ = user.Insert(ctx, database.DB, boil.Infer())
	// Test login
	test := assert.New(t)
	testCases := []struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{
		{
			UserName: "test",
			Password: "test",
		},
		{
			UserName: "test",
			Password: "test",
		},
		{
			UserName: "test2",
			Password: "test",
		},
		{
			UserName: "test",
			Password: "test2",
		},
	}
	expectedStatusCodes := []int{401, 200, 401, 401}
	for index, testCase := range testCases {
		if index == 1 {
			// Activate test user
			user.Active = true
			_, _ = user.Update(ctx, database.DB, boil.Infer())
		}

		app := setup()
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)
		test.Equal("application/json", response.Header.Get("Content-Type"))
		if response.StatusCode == 200 {
			var result responses.TokenResponse
			_ = json.NewDecoder(response.Body).Decode(&result)
			test.NotEmpty(result.AccessToken)
			test.NotEmpty(result.RefreshToken)
		} else {
			var result responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&result)
			test.NotEmpty(result.Message)
			test.Equal("INVALID_LOGIN_CREDENTIALS", result.Message)
		}
	}
	tearDownDatabase()
}

func TestLogout(t *testing.T) {
	setupDatabase()
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
	// Test logout
	test := assert.New(t)
	testCase := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{
		UserName: "test",
		Password: "test",
	}
	// Login
	app := setup()
	body, _ := json.Marshal(testCase)
	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var result responses.TokenResponse
	_ = json.NewDecoder(response.Body).Decode(&result)
	// Logout
	req = httptest.NewRequest("POST", "/api/v1/auth/logout", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))
	response, _ = app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var successRes responses.SuccessResponse
	_ = json.NewDecoder(response.Body).Decode(&successRes)
	test.NotEmpty(successRes.Message)
	test.Equal("LOGGED_OUT", successRes.Message)
	// Logout again
	req = httptest.NewRequest("POST", "/api/v1/auth/logout", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))
	response, _ = app.Test(req, 5000)
	test.Equal(401, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var errorRes responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&errorRes)
	test.NotEmpty(errorRes.Message)
	test.Equal("UNAUTHORIZED", errorRes.Message)
	tearDownDatabase()
}

func TestRefresh(t *testing.T) {
	setupDatabase()
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
	// Test logout
	test := assert.New(t)
	testCase := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{
		UserName: "test",
		Password: "test",
	}
	// Login
	app := setup()
	body, _ := json.Marshal(testCase)
	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var result responses.TokenResponse
	_ = json.NewDecoder(response.Body).Decode(&result)
	// Refresh
	tokenRequest := struct {
		RefreshToken string `json:"refreshToken"`
	}{
		RefreshToken: result.RefreshToken,
	}
	body, _ = json.Marshal(tokenRequest)
	req = httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", result.AccessToken))
	response, _ = app.Test(req, 5000)
	test.Equal(201, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var tokenRes responses.TokenResponse
	_ = json.NewDecoder(response.Body).Decode(&tokenRes)
	test.NotEmpty(tokenRes.AccessToken)
	test.NotEmpty(tokenRes.RefreshToken)
	tearDownDatabase()
}
