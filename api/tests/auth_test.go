package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
		response, err := app.Test(req)
		if err != nil {
			panic(err)
		}

		test.Equal(expectedStatusCodes[index], response.StatusCode)
	}
	tearDownDatabase()
}
