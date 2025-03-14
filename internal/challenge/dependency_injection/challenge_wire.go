package dependency_injection

import (
	"database/sql"

	"github.com/ahmadammarm/scrolless-backend/internal/challenge/handler"
	"github.com/ahmadammarm/scrolless-backend/internal/challenge/repository"
	"github.com/ahmadammarm/scrolless-backend/internal/challenge/service"
	userRepo "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
	"github.com/go-playground/validator/v10"
)

func InitializedChallengeService(db *sql.DB, validator *validator.Validate) *handler.ChallengeHandler {
    challengeRepo := repository.NewChallengeRepository(db)
    userRepository := userRepo.NewUserRepository(db)
    challengeService := service.NewChallengeService(challengeRepo, userRepository)
    challengeHandler := handler.NewChallengeHandler(challengeService, validator)
    return challengeHandler
}