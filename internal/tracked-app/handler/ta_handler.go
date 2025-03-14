package handler

import (
	"strconv"

	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/entity"
	trackedAppService "github.com/ahmadammarm/scrolless-backend/internal/tracked-app/service"
	"github.com/ahmadammarm/scrolless-backend/utils/form"
	"github.com/ahmadammarm/scrolless-backend/utils/middleware"
	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TrackedAppHandler struct {
	trackedAppService trackedAppService.TrackedAppService
	validation        *validator.Validate
}

func (handler *TrackedAppHandler) ListTrackedApp(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	apps, err := handler.trackedAppService.ListTrackedApp(userId)
	if err != nil {
		return response.JSON(context, 500, "List Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "List Tracked App Success", apps)
}

func (handler *TrackedAppHandler) GetTrackedAppByID(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	trackedAppId, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	trackedApp, err := handler.trackedAppService.GetTrackedAppByID(userId, trackedAppId)

	if err != nil {
		return response.JSON(context, 500, "Get Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "Get Tracked App Success", trackedApp)
}

func (handler *TrackedAppHandler) CreateTrackedApp(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	trackedApp := new(entity.TrackedAppsRequest)
	if err := context.BodyParser(trackedApp); err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	if err := handler.validation.Struct(trackedApp); err != nil {
		errorValidations := form.ErrorFormValidation(err)
		return response.JSON(context, 400, "Invalid Request", errorValidations)
	}

	if err := handler.trackedAppService.CreateTrackedApp(userId, trackedApp); err != nil {
		return response.JSON(context, 500, "Create Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "Create Tracked App Success", nil)
}

func (handler *TrackedAppHandler) DeleteTrackedApp(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	trackedAppId, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	err = handler.trackedAppService.DeleteTrackedApp(userId, trackedAppId)
	if err != nil {
		return response.JSON(context, 500, "Delete Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "Delete Tracked App Success", nil)
}

func (handler *TrackedAppHandler) ActivateTrackedApp(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	trackedAppId, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	err = handler.trackedAppService.ActivateTrackedApp(userId, trackedAppId)
	if err != nil {
		return response.JSON(context, 500, "Activate Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "Activate Tracked App Success", nil)
}

func (handler *TrackedAppHandler) DeactivateTrackedApp(context *fiber.Ctx) error {
	userId, ok := context.Locals("user_id").(int)
	if !ok {
		return response.JSON(context, 401, "Unauthorized", nil)
	}

	trackedAppId, err := strconv.Atoi(context.Params("id"))

	if err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	err = handler.trackedAppService.DeactivateTrackedApp(userId, trackedAppId)
	if err != nil {
		return response.JSON(context, 500, "Deactivate Tracked App Failed", nil)
	}

	return response.JSON(context, 200, "Deactivate Tracked App Success", nil)
}

func (handler *TrackedAppHandler) Router(router fiber.Router) {
	router.Use(middleware.ProtectedJWT())
	router.Get("/tracked-apps", handler.ListTrackedApp)
	router.Get("/tracked-apps/:id", handler.GetTrackedAppByID)
	router.Post("/tracked-apps", handler.CreateTrackedApp)
	router.Delete("/tracked-apps/:id", handler.DeleteTrackedApp)
	router.Post("/tracked-apps/:id/activate", handler.ActivateTrackedApp)
	router.Post("/tracked-apps/:id/deactivate", handler.DeactivateTrackedApp)
}

func NewTrackedAppHandler(trackedAppService trackedAppService.TrackedAppService, validator *validator.Validate) *TrackedAppHandler {
	return &TrackedAppHandler{
		trackedAppService: trackedAppService,
		validation:        validator,
	}
}
