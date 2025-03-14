package handler

import (
	"strconv"

	"github.com/ahmadammarm/scrolless-backend/internal/challenge/entity"
	challengeService "github.com/ahmadammarm/scrolless-backend/internal/challenge/service"
	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ChallengeHandler struct {
	challengeService challengeService.ChallengeService
	validation       *validator.Validate
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

func (handler *ChallengeHandler) ListChallenge(context *fiber.Ctx) error {
	userID, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	challenges, err := handler.challengeService.ListChallenge(userID)
	if err != nil {
		return response.JSON(context, 500, "List Challenge Failed", nil)
	}

	return response.JSON(context, 200, "List Challenge Success", challenges)
}

func (handler *ChallengeHandler) AddPointsByChallengeDone(context *fiber.Ctx) error {
	userID, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	challengeID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	err = handler.challengeService.AddPointsByChallengeDone(userID, challengeID)
	if err != nil {
		return response.JSON(context, 500, "Add Points By Challenge Done Failed", nil)
	}

	return response.JSON(context, 200, "Add Points By Challenge Done Success", nil)
}

func (handler *ChallengeHandler) GetChallengeByID(context *fiber.Ctx) error {
	userID, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	challengeID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	challenge, err := handler.challengeService.GetChallengeByID(userID, challengeID)
	if err != nil {
		return response.JSON(context, 500, "Get Challenge Failed", nil)
	}

	return response.JSON(context, 200, "Get Challenge Success", challenge)
}

func (handler *ChallengeHandler) DeleteChallenge(context *fiber.Ctx) error {
	userID, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	challengeID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	err = handler.challengeService.DeleteChallenge(userID, challengeID)
	if err != nil {
		return response.JSON(context, 500, "Delete Challenge Failed", nil)
	}

	return response.JSON(context, 200, "Delete Challenge Success", nil)
}

func (handler *ChallengeHandler) Router(router fiber.Router) {
	router.Post("/challenge", handler.CreateChallenge)
	router.Get("/challenge", handler.ListChallenge)
	router.Get("/challenge/:id", handler.GetChallengeByID)
	router.Delete("/challenge/:id", handler.DeleteChallenge)
	router.Post("/challenge/:id/done", handler.AddPointsByChallengeDone)
}

func NewChallengeHandler(challengeService challengeService.ChallengeService, validator *validator.Validate) *ChallengeHandler {
    return &ChallengeHandler{
        challengeService: challengeService,
        validation:       validator,
    }
}
