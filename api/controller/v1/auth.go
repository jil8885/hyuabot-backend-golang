package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/requests"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
	"github.com/hyuabot-developers/hyuabot-backend-golang/utils"
)

func SignUp(c *fiber.Ctx) error {
	var request requests.SignUpRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	} else if request.Password == "" || request.Username == "" || request.Name == "" || request.Email == "" || request.Phone == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	exists, err := models.AdminUserExists(c.Context(), database.DB, request.Username)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if exists {
		return c.Status(http.StatusConflict).JSON(responses.ErrorResponse{Message: "USER_ALREADY_EXISTS"})
	}

	user := models.AdminUser{
		UserID:   request.Username,
		Password: hashedPassword,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Active:   false,
	}
	err = user.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	return c.Status(http.StatusCreated).JSON(responses.SuccessResponse{Message: "USER_CREATED"})
}

func Login(c *fiber.Ctx) error {
	var request requests.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	exists, err := models.AdminUserExists(c.Context(), database.DB, request.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if !exists {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "INVALID_LOGIN_CREDENTIALS"})
	}

	user, err := models.FindAdminUser(c.Context(), database.DB, request.Username)
	if err != nil {
		fmt.Printf("%s", err)
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if !utils.CheckPasswordHash(request.Password, user.Password) || !user.Active {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "INVALID_LOGIN_CREDENTIALS"})
	}

	token, err := utils.CreateJWTToken(request.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	err = utils.CreateAuth(request.Username, token)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	return c.Status(http.StatusOK).JSON(responses.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
}

func Logout(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetaData(c.Request())
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "UNAUTHORIZED"})
	}

	deleted, err := utils.DeleteAuth(au.AccessUUID)
	if err != nil || deleted == 0 {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "UNAUTHORIZED"})
	}

	return c.Status(http.StatusOK).JSON(responses.SuccessResponse{Message: "LOGGED_OUT"})
}

func Refresh(c *fiber.Ctx) error {
	mapToken := map[string]string{}
	if err := c.BodyParser(&mapToken); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	refreshToken := mapToken["refreshToken"]
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED_SIGNING_METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "REFRESH_TOKEN_EXPIRED"})
	}

	if !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
		}
		userID := claims["user_id"].(string)
		deleted, delErr := utils.DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 {
			return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "UNAUTHORIZED"})
		}
		ts, createErr := utils.CreateJWTToken(userID)
		if createErr != nil {
			return c.Status(http.StatusForbidden).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		}
		saveErr := utils.CreateAuth(userID, ts)
		if saveErr != nil {
			return c.Status(http.StatusForbidden).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		}
		return c.Status(http.StatusCreated).JSON(responses.TokenResponse{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		})
	}
	return c.Status(http.StatusUnauthorized).JSON(responses.ErrorResponse{Message: "REFRESH_TOKEN_EXPIRED"})
}
