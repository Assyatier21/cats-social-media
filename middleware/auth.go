package middleware

import (
	"errors"
	"fmt"
	"log"
	"time"

	cfg "github.com/backend-magang/cats-social-media/config"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	Authorization = "Authorization"
)

func getJWTSecretKey() string {
	return cfg.Load().JWTSecret
}

func ClaimToken(ctx echo.Context) entity.UserClaimsResponse {
	return ctx.Get("user").(entity.UserClaimsResponse)
}

func GenerateToken(registry entity.User) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = registry.Name
	claims["email"] = registry.Email
	claims["expired_at"] = time.Now().Add(8 * time.Hour)

	t, err = token.SignedString([]byte(getJWTSecretKey()))
	if err != nil {
		log.Println("[Middleware] failed to signed jwt token, err: ", err)
		return
	}

	return
}

func ParseTokenJWT(tokenString string) (userClaims entity.UserClaimsResponse, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(getJWTSecretKey()), nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userClaims, errors.New("invalid token claims")
	}

	expiredAtStr, _ := claims["expired_at"].(string)
	expiredAt, err := time.Parse(time.RFC3339, expiredAtStr)
	if err != nil {
		return userClaims, fmt.Errorf("failed to parse expired_at value: %v", err)
	}
	userClaims = entity.UserClaimsResponse{
		Name:      claims["name"].(string),
		Email:     claims["email"].(string),
		ExpiredAt: expiredAt,
	}

	return
}

func IsPrivate(role string, private bool) bool {
	return (role != "admin" && private)
}

func IsTokenExpired(t time.Time) bool {
	now := time.Now()
	return t.Before(now)
}
