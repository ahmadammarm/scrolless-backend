package service

import (
	"fmt"
	"github.com/ahmadammarm/scrolless-backend/internal/challenge/entity"
	challengeRepo "github.com/ahmadammarm/scrolless-backend/internal/challenge/repository"
	userRepo "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
)

type ChallengeService interface {
	CreateChallenge(userID int, challenge *entity.Challenge) error
	ListChallenge(userID int) (*entity.ChallengeListResponse, error)
	AddPointsByChallengeDone(userID, challengeID int) error
	GetChallengeByID(userID, challengeID int) (*entity.ChallengeResponse, error)
	DeleteChallenge(userID, challengeID int) error
}

type challengeService struct {
	challengeRepo challengeRepo.ChallengeRepository
	userRepo      userRepo.UserRepository
}

func (service *challengeService) CreateChallenge(userID int, challenge *entity.Challenge) error {
	userExists, err := service.userRepo.IsUserExists(userID)

	if err != nil {
		return fmt.Errorf("failed to check user: %w", err)
	}

	if !userExists {
		return fmt.Errorf("user not found or not logged in")
	}

	challenge.UserID = userID
	err = service.challengeRepo.CreateChallenge(challenge)

	if err != nil {
		return fmt.Errorf("failed to create challenge: %w", err)
	}

	return nil
}

func (service *challengeService) ListChallenge(userID int) (*entity.ChallengeListResponse, error) {
    userExists, err := service.userRepo.IsUserExists(userID)

    if err != nil {
        return nil, fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return nil, fmt.Errorf("user not found or not logged in")
    }

    challenges, err := service.challengeRepo.ListChallenge()

    if err != nil {
        return nil, fmt.Errorf("failed to get challenges: %w", err)
    }

    return challenges, nil
}

func (service *challengeService) AddPointsByChallengeDone(userID, challengeID int) error {
    userExists, err := service.userRepo.IsUserExists(userID)

    if err != nil {
        return fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return fmt.Errorf("user not found or not logged in")
    }

    err = service.challengeRepo.AddPointsByChallengeDone(challengeID)

    if err != nil {
        return fmt.Errorf("failed to add points: %w", err)
    }

    return nil
}

func (service *challengeService) GetChallengeByID(userID, challengeID int) (*entity.ChallengeResponse, error) {
    userExists, err := service.userRepo.IsUserExists(userID)

    if err != nil {
        return nil, fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return nil, fmt.Errorf("user not found or not logged in")
    }

    challenge, err := service.challengeRepo.GetChallengeByID(challengeID)

    if err != nil {
        return nil, fmt.Errorf("failed to get challenge: %w", err)
    }

    return challenge, nil
}

func (service *challengeService) DeleteChallenge(userID, challengeID int) error {
    userExists, err := service.userRepo.IsUserExists(userID)

    if err != nil {
        return fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return fmt.Errorf("user not found or not logged in")
    }

    err = service.challengeRepo.DeleteChallenge(challengeID)

    if err != nil {
        return fmt.Errorf("failed to delete challenge: %w", err)
    }

    return nil
}

func NewChallengeService(challengeRepo challengeRepo.ChallengeRepository, userRepo userRepo.UserRepository) ChallengeService {
    return &challengeService{challengeRepo, userRepo}
}

