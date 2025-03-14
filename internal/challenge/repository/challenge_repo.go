package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmadammarm/scrolless-backend/internal/challenge/entity"
)

type ChallengeRepository interface {
	CreateChallenge(challenge *entity.Challenge) error
	ListChallenge() (*entity.ChallengeListResponse, error)
	AddPointsByChallengeDone(challengeID int) error
	GetChallengeByID(challengeID int) (*entity.ChallengeResponse, error)
	DeleteChallenge(challengeID int) error
}

type challengeRepository struct {
	db *sql.DB
}

func (repo *challengeRepository) CreateChallenge(challenge *entity.Challenge) error {
	query := `INSERT INTO challenges (title, user_id, description) VALUES ($1, $2, $3)`
	_, err := repo.db.Exec(query, challenge.Title, challenge.UserID, challenge.Description)
	if err != nil {
		return err
	}

	return nil
}

func (repo *challengeRepository) ListChallenge() (*entity.ChallengeListResponse, error) {
	query := `SELECT id, title, description, status, points FROM challenges`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var challenges []entity.ChallengeResponse
	for rows.Next() {
		var challenge entity.ChallengeResponse
		err := rows.Scan(&challenge.ID, &challenge.Title, &challenge.Description, &challenge.Status, &challenge.Points)
		if err != nil {
			return nil, err
		}

		challenges = append(challenges, challenge)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &entity.ChallengeListResponse{Challenges: challenges}, nil
}

func (repo *challengeRepository) AddPointsByChallengeDone(challengeID int) error {
	query := `UPDATE challenges SET status = 'sudah', points = points + 10 WHERE id = $1`
	_, err := repo.db.Exec(query, challengeID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *challengeRepository) GetChallengeByID(challengeID int) (*entity.ChallengeResponse, error) {
	query := `SELECT id, title, description, status, points FROM challenges WHERE id = $1`
	challenge := &entity.ChallengeResponse{}

	err := repo.db.QueryRow(query, challengeID).Scan(&challenge.ID, &challenge.Title, &challenge.Description, &challenge.Status, &challenge.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("challenge not found")
		}
	}

	return challenge, nil
}

func (repo *challengeRepository) DeleteChallenge(challengeID int) error {
	query := `DELETE FROM challenges WHERE id = $1`
	_, err := repo.db.Exec(query, challengeID)
	if err != nil {
		return err
	}

	return nil
}

func NewChallengeRepository(db *sql.DB) ChallengeRepository {
	return &challengeRepository{db}
}
