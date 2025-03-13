package service

import (
	"errors"
	"fmt"

	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/entity"
	trackedAppRepo "github.com/ahmadammarm/scrolless-backend/internal/tracked-app/repository"
	userRepo "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
)

type TrackedAppService interface {
	ListTrackedApp(userID int) (*entity.TrackedAppsListResponse, error)
	GetTrackedAppByID(userID, trackedAppID int) (*entity.TrackedAppsResponse, error)
	CreateTrackedApp(userID int, trackedApp *entity.TrackedAppsRequest) error
	DeleteTrackedApp(userID, trackedAppID int) error
}

type trackedAppService struct {
	trackedAppRepo trackedAppRepo.TrackedAppRepository
	userRepo       userRepo.UserRepository
}

func (service *trackedAppService) ListTrackedApp(userID int) (*entity.TrackedAppsListResponse, error) {
	userExists, err := service.userRepo.IsUserExists(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user: %w", err)
	}
	if !userExists {
		return nil, errors.New("user not found or not logged in")
	}

	apps, err := service.trackedAppRepo.ListTrackedApp()
	if err != nil {
		return nil, fmt.Errorf("failed to get tracked apps: %w", err)
	}

	filteredApps := []entity.TrackedApps{}
	for _, app := range apps.Apps {
		if app.UserID == userID {
			filteredApps = append(filteredApps, app)
		}
	}

	return &entity.TrackedAppsListResponse{Apps: filteredApps}, nil
}

func (service *trackedAppService) GetTrackedAppByID(userID, trackedAppID int) (*entity.TrackedAppsResponse, error) {
	userExists, err := service.userRepo.IsUserExists(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return nil, errors.New("user not found or not logged in")
    }

    app, err := service.trackedAppRepo.GetTrackedAppByID(trackedAppID)

    if err != nil {
        return nil, fmt.Errorf("failed to get tracked app: %w", err)
    }

    if app.UserID != userID {
        return nil, errors.New("access denied: app is not owned by user")
    }

    return app, nil
}

func (service *trackedAppService) CreateTrackedApp(userID int, trackedApp *entity.TrackedAppsRequest) error {
    userExists, err := service.userRepo.IsUserExists(userID)
    if err != nil {
        return fmt.Errorf("failed to check user: %w", err)
    }
    if !userExists {
        return errors.New("user not found or not logged in")
    }

    trackedApp.UserID = userID
    err = service.trackedAppRepo.CreateTrackedApp(trackedApp)
    if err != nil {
        return fmt.Errorf("failed to create tracked app: %w", err)
    }
    return nil
}

func (service *trackedAppService) DeleteTrackedApp(userID, trackedAppID int) error {
	userExists, err := service.userRepo.IsUserExists(userID)
    if err != nil {
        return fmt.Errorf("failed to check user: %w", err)
    }

    if !userExists {
        return errors.New("user not found or not logged in")
    }

    app, err := service.trackedAppRepo.GetTrackedAppByID(trackedAppID)

    if err != nil {
        return fmt.Errorf("failed to get tracked app: %w", err)
    }

    if app.UserID != userID {
        return errors.New("access denied: app is not owned by user")
    }

    return service.trackedAppRepo.DeleteTrackedApp(trackedAppID)
}

func NewTrackedAppService(trackedAppRepo trackedAppRepo.TrackedAppRepository, userRepo userRepo.UserRepository) TrackedAppService {
    return &trackedAppService{trackedAppRepo: trackedAppRepo, userRepo: userRepo}
}