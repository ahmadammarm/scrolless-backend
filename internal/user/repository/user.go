package repository

import (
	"database/sql"
	"errors"
	"time"

	userEntity "github.com/ahmadammarm/scrolless-backend/internal/user/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	ListUser() (*userEntity.UserListResponse, error)
	GetUserByID(userId int) (*userEntity.UserDetailResponse, error)
	RegisterUser(user *userEntity.UserRegister) error
	LoginUser(user *userEntity.UserLogin) (*userEntity.UserJWT, error)
}

type userRepository struct {
	db *sql.DB
}

func (repo *userRepository) ListUser() (*userEntity.UserListResponse, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []userEntity.UserListSubResponse
	for rows.Next() {
		var user userEntity.UserListSubResponse
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &userEntity.UserListResponse{Users: users}, nil
}

func (repo *userRepository) GetUserByID(userId int) (*userEntity.UserDetailResponse, error) {
    query := `SELECT id, name, email FROM users WHERE id = $1`
    user := &userEntity.UserDetailResponse{}

    err := repo.db.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return user, nil
}

func (repo *userRepository) RegisterUser(user *userEntity.UserRegister) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	err = tx.QueryRow(query, user.Name, user.Email, hashedPassword).Scan(&user.ID)
	if err != nil {
		return err
	}

    return nil
}

func (repo *userRepository) LoginUser(user *userEntity.UserLogin) (*userEntity.UserJWT, error) {
	query := `SELECT id, name, email, password, role FROM users WHERE email = $1`
	jwtUser := &userEntity.UserJWT{}

	err := repo.db.QueryRow(query, user.Email).Scan(&jwtUser.ID, &jwtUser.Name, &jwtUser.Email, &jwtUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(jwtUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return jwtUser, nil

}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db}
}

