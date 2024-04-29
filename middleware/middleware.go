package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/backend-magang/cats-social-media/models"
	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func TokenValidationMiddleware(private bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var userClaims = entity.UserClaimsResponse{}
			tokenString := c.Request().Header.Get(Authorization)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

			userClaims.Name = claims["name"].(string)
			userClaims.Email = claims["email"].(string)
			expiredAtStr, _ := claims["expired_at"].(string)
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
