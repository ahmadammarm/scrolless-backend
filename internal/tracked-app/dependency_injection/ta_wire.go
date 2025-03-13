package dependency_injection

import (
	"database/sql"
	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/handler"
	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/repository"
	userRepository "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/service"
	"github.com/go-playground/validator/v10"
)

func InitializedTrackedAppService(db *sql.DB, validator *validator.Validate) *handler.TrackedAppHandler {
    trackedAppRepo := repository.NewTrackedAppRepository(db)
    userRepo := userRepository.NewUserRepository(db)
    trackedAppService := service.NewTrackedAppService(trackedAppRepo, userRepo)
    trackedAppHandler := handler.NewTrackedAppHandler(trackedAppService, validator)
    return trackedAppHandler
}