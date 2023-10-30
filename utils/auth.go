package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/twinj/uuid"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"

	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
)

func CreateJWTToken(username string) (*responses.TokenDetails, error) {
	details := &responses.TokenDetails{}
	details.AccessTokenExp = time.Now().Add(time.Minute * 15).Unix()
	details.AccessUUID = uuid.NewV4().String()
	details.RefreshTokenExp = time.Now().Add(time.Hour * 24 * 7).Unix()
	details.RefreshUUID = uuid.NewV4().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = details.AccessUUID
	atClaims["user_id"] = username
	atClaims["exp"] = details.AccessTokenExp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	details.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = details.RefreshUUID
	rtClaims["user_id"] = username
	rtClaims["exp"] = details.RefreshTokenExp
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	details.RefreshToken, err = rt.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return details, nil
}

func CreateAuth(username string, td *responses.TokenDetails) error {
	at := time.Unix(td.AccessTokenExp, 0)
	rt := time.Unix(td.RefreshTokenExp, 0)
	now := time.Now()

	errAccess := database.Client.Set(td.AccessUUID, username, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := database.Client.Set(td.RefreshUUID, username, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractToken(r *fasthttp.Request) string {
	bearerToken := r.Header.Peek("Authorization")
	strArr := strings.Split(string(bearerToken), " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *fasthttp.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED_SIGNING_METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CheckTokenIsValid(r *fasthttp.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetaData(r *fasthttp.Request) (*responses.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		return &responses.AccessDetails{
			AccessUUID: accessUUID,
			UserName:   userID,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *responses.AccessDetails) (string, error) {
	userID, err := database.Client.Get(authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}

func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := database.Client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
