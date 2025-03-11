package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(context *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := Response{
		Message: message,
		Data:    data,
	}
	return context.Status(statusCode).JSON(response)
}
