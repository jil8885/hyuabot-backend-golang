package tests

import (
	"bytes"
	"context"
	"encoding/json"
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
		body, err := json.Marshal(testCase)
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response, err := app.Test(req, 5000)
		if err != nil {
			panic(err)
		}

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
	err := user.Insert(ctx, database.DB, boil.Infer())
	if err != nil {
		panic(err)
	}
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
			_, err := user.Update(ctx, database.DB, boil.Infer())
			if err != nil {
				panic(err)
			}
		}

		app := setup()
		body, err := json.Marshal(testCase)
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		response, err := app.Test(req, 5000)
		if err != nil {
			panic(err)
		}

		test.Equal(expectedStatusCodes[index], response.StatusCode)
		test.Equal("application/json", response.Header.Get("Content-Type"))
		if response.StatusCode == 200 {
			var result responses.TokenResponse
			err := json.NewDecoder(response.Body).Decode(&result)
			if err != nil {
				panic(err)
			}
			test.NotEmpty(result.AccessToken)
			test.NotEmpty(result.RefreshToken)
		} else {
			var result responses.ErrorResponse
			err := json.NewDecoder(response.Body).Decode(&result)
			if err != nil {
				panic(err)
			}
			test.NotEmpty(result.Message)
			test.Equal("INVALID_LOGIN_CREDENTIALS", result.Message)
		}
	}
	tearDownDatabase()
}
