package main

import (
	"log"
	"os"

	"github.com/ahmadammarm/scrolless-backend/config"
	challenges "github.com/ahmadammarm/scrolless-backend/internal/challenge/dependency_injection"
	trackedApps "github.com/ahmadammarm/scrolless-backend/internal/tracked-app/dependency_injection"
	users "github.com/ahmadammarm/scrolless-backend/internal/user/dependency_injection"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	db, error := config.PostgresInit()

	if error != nil {
		log.Printf("Error connecting to Postgres: %v", error)
		os.Exit(1)
	}

	defer db.Close()

	application := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := application.Group("/api")

	users.InitializedUserService(db, validator.New()).Router(api)
	trackedApps.InitializedTrackedAppService(db, validator.New()).Router(api)
	challenges.InitializedChallengeService(db, validator.New()).Router(api)

	if error := application.Listen(":3000"); error != nil {
		log.Printf("Error starting server: %v", error)
	}
}
