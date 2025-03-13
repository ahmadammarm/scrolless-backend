package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ahmadammarm/scrolless-backend/internal/tracked-app/entity"
)

type TrackedAppRepository interface {
	CreateTrackedApp(trackedApp *entity.TrackedAppsRequest) error
	ListTrackedApp() (*entity.TrackedAppsListResponse, error)
	GetTrackedAppByID(trackedAppId int) (*entity.TrackedAppsResponse, error)
	DeleteTrackedApp(trackedAppId int) error
	ActivateTrackedApp(trackedAppId int) error
    DeactivateTrackedApp(trackedAppId int) error
}

type trackedAppRepository struct {
	db *sql.DB
}

func (repo *trackedAppRepository) ListTrackedApp() (*entity.TrackedAppsListResponse, error) {
	query := `SELECT id, user_id, app_name, status, created_at FROM tracked_apps`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var apps []entity.TrackedApps
	for rows.Next() {
		var app entity.TrackedApps
		err := rows.Scan(&app.ID, &app.UserID, &app.AppName, &app.Status, &app.CreatedAt)
		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &entity.TrackedAppsListResponse{Apps: apps}, nil
}

func (repo *trackedAppRepository) GetTrackedAppByID(trackedAppId int) (*entity.TrackedAppsResponse, error) {
	query := `SELECT id, user_id, app_name, status, created_at FROM tracked_apps WHERE id = $1`
	app := &entity.TrackedAppsResponse{}

	err := repo.db.QueryRow(query, trackedAppId).Scan(&app.ID, &app.UserID, &app.AppName, &app.Status, &app.CreatedAt)
	if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("tracked app not found")
        }
    }

	return app, nil
}


func (repo *trackedAppRepository) CreateTrackedApp(trackedApp *entity.TrackedAppsRequest) error {
    query := `INSERT INTO tracked_apps (user_id, app_name, status, duration, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

    trackedApp.CreatedAt = time.Now().Unix()

    err := repo.db.QueryRow(query, trackedApp.UserID, trackedApp.AppName, trackedApp.Status, trackedApp.CreatedAt, trackedApp.Duration).Scan(&trackedApp.ID)
    if err != nil {
        return err
    }

    return nil
}


func (repo *trackedAppRepository) DeleteTrackedApp(trackedAppId int) error {
	query := `DELETE FROM tracked_apps WHERE id = $1`
	_, err := repo.db.Exec(query, trackedAppId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *trackedAppRepository) ActivateTrackedApp(trackedAppId int) error {
    query := `UPDATE tracked_apps SET status = true, duration = $1, end_time = $2 WHERE id = $3`
    _, err := repo.db.Exec(query, trackedAppId)
    if err != nil {
        return err
    }
    return nil
}

func (repo *trackedAppRepository) DeactivateTrackedApp(trackedAppId int) error {
    query := `UPDATE tracked_apps SET status = false, duration = 0, end_time = 0 WHERE id = $1`
    _, err := repo.db.Exec(query, trackedAppId)
    if err != nil {
        return err
    }

    return nil
}

func NewTrackedAppRepository(db *sql.DB) TrackedAppRepository {
    return &trackedAppRepository{db}
}
