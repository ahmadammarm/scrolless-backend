package handler

import (
	"strconv"

	"github.com/ahmadammarm/scrolless-backend/internal/user/entity"
	userService "github.com/ahmadammarm/scrolless-backend/internal/user/service"
	"github.com/ahmadammarm/scrolless-backend/utils/form"
	"github.com/ahmadammarm/scrolless-backend/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService userService.UserService
	validation  *validator.Validate
}

func (handler *userHandler) RegisterUser(context *fiber.Ctx) error {
	user := new(entity.UserRegister)
	if err := context.BodyParser(user); err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	if err := handler.validation.Struct(user); err != nil {
		errorValidations := form.ErrorFormValidation(err)
		return response.JSON(context, 400, "Invalid Request", errorValidations)
	}

	if err := handler.userService.RegisterUser(user); err != nil {
		return response.JSON(context, 500, "Register User Failed", nil)
	}

	return response.JSON(context, 200, "Register User Success", nil)
}

func (handler *userHandler) LoginUser(context *fiber.Ctx) error {

	loginReq := new(entity.UserLogin)

	if err := context.BodyParser(loginReq); err != nil {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	token, err := handler.userService.LoginUser(loginReq)

	if err != nil {
		return response.JSON(context, 401, "Login User Failed", nil)
	}

	return response.JSON(context, 200, "Login User Success", token)
}

func (handler *userHandler) ListUser(context *fiber.Ctx) error {
	users, err := handler.userService.ListUser()
	if err != nil {
		return response.JSON(context, 500, "List User Failed", nil)
	}

	return response.JSON(context, 200, "List User Success", users)
}

func (handler *userHandler) GetUserByID(context *fiber.Ctx) error {
	userIdString := context.Params("id")
	userId, err := strconv.Atoi(userIdString)

	if err != nil || userId < 1 {
		return response.JSON(context, 400, "Invalid Request", nil)
	}

	user, err := handler.userService.GetUserByID(userId)
	if err != nil {
		return response.JSON(context, 404, "User Not Found", nil)
	}

	return response.JSON(context, 200, "Get User Success", user)
}

func (handler *userHandler) Router(router fiber.Router) {
	router.Post("/user/register", handler.RegisterUser)
	router.Post("/user/login", handler.LoginUser)
	router.Get("/users", handler.ListUser)
	router.Get("/user/:id", handler.GetUserByID)
}

func NewUserHandler(userService userService.UserService, validation *validator.Validate) *userHandler {
	return &userHandler{
		userService: userService,
		validation:  validation,
	}
}
