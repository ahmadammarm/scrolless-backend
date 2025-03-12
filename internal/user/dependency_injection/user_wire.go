// dependency injection digunakan untuk menginisialisasi service dan repository yang akan digunakan
// pada handler. Dengan menggunakan dependency injection, kita bisa mengganti service dan repository

package dependency_injection

import (
    "database/sql"
    "github.com/ahmadammarm/scrolless-backend/internal/user/handler"
    "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
    "github.com/ahmadammarm/scrolless-backend/internal/user/service"
    "github.com/go-playground/validator/v10"
)

func InitializedUserService(db *sql.DB, validator *validator.Validate) *handler.UserHandler {
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService, validator)
    return userHandler
}