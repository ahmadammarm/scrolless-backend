package middleware

import (
	"os"
	"strings"

	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ProtectedJWT() fiber.Handler {
	return func(context *fiber.Ctx) error {
		stringToken := context.Get("Authorization")
		if stringToken == "" {
			return response.JSON(context, 401, "Unauthorized", nil)
		}

		stringToken = strings.TrimPrefix(stringToken, "Bearer ")

		token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return response.JSON(context, 401, "Unauthorized", nil)
		}

		context.Locals("user", token)

		return context.Next()
	}
}
