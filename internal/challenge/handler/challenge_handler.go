package handler

import (
	"github.com/ahmadammarm/scrolless-backend/internal/challenge/entity"
	challengeService "github.com/ahmadammarm/scrolless-backend/internal/challenge/service"
	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ChallengeHandler struct {
    challengeService challengeService.ChallengeService
    validation *validator.Validate
}

func (handler *ChallengeHandler) CreateChallenge(context *fiber.Ctx) error {
    userID, ok := context.Locals("user_id").(int)
    if !ok {
        return response.JSON(context, 401, "Unauthorized", nil)
    }

    challenge := new(entity.Challenge)
    if err := context.BodyParser(challenge); err != nil {
        return response.JSON(context, 400, "Invalid Request", nil)
    }

    err := handler.challengeService.CreateChallenge(userID, challenge)
    if err != nil {
        return response.JSON(context, 500, "Create Challenge Failed", nil)
    }

    return response.JSON(context, 200, "Create Challenge Success", nil)
}

