package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func TokenValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var userClaims = entity.UserClaimsResponse{}
			tokenString := c.Request().Header.Get(Authorization)

			if !strings.Contains(tokenString, "Bearer ") {
				response := models.StandardResponse{
					Code:    http.StatusUnauthorized,
					Status:  constant.FAILED,
					Message: "Invalid token",
				}
				return c.JSON(http.StatusUnauthorized, response)
			}

			tokenJWT := strings.Split(tokenString, " ")[1]

			token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid token signing method")
				}
				return []byte(getJWTSecretKey()), nil
			})
			if err != nil {
				response := models.StandardResponse{
					Code:    http.StatusUnauthorized,
					Status:  constant.FAILED,
					Message: "Invalid signing token method",
					Error:   err,
				}
				return c.JSON(http.StatusUnauthorized, response)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				response := models.StandardResponse{
					Code:    http.StatusUnauthorized,
					Status:  constant.FAILED,
					Message: "Invalid token claims",
					Error:   errors.New("invalid token claims"),
				}
				return c.JSON(http.StatusUnauthorized, response)
			}

			userClaims.ID = cast.ToInt(claims["id"])
			userClaims.Name = cast.ToString(claims["name"])
			userClaims.Email = cast.ToString(claims["email"])
			expiredAtStr := cast.ToString(claims["expired_at"])
			expiredAt, _ := time.Parse(time.RFC3339, expiredAtStr)
			userClaims.ExpiredAt = expiredAt

			if IsTokenExpired(expiredAt) {
				response := models.StandardResponse{
					Code:    http.StatusUnauthorized,
					Status:  constant.FAILED,
					Message: "Token already expired",
					Error:   errors.New("token already expired"),
				}
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("user", userClaims)
			return next(c)
		}
	}
}
