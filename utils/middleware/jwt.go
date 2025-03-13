package middleware

import (
	"os"
	"strings"
	"fmt"

	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
)

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func ProtectedJWT() fiber.Handler {
	return func(context *fiber.Ctx) error {
		stringToken := context.Get("Authorization")
		if stringToken == "" {
			return response.JSON(context, 401, "Unauthorized: No Token Provided", nil)
		}


		stringToken = strings.TrimPrefix(stringToken, "Bearer ")

		token, err := jwt.ParseWithClaims(stringToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				return nil, fmt.Errorf("JWT_SECRET is missing")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return response.JSON(context, 401, "Unauthorized: Invalid Token", nil)
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return response.JSON(context, 401, "Unauthorized: Token Claims Invalid", nil)
		}

		context.Locals("user_id", claims.UserID)

		return context.Next()
	}
}

