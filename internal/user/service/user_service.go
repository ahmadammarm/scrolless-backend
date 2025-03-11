package service

import (
	"errors"
	"os"
	"time"

	userEntity "github.com/ahmadammarm/scrolless-backend/internal/user/entity"
	userRepo "github.com/ahmadammarm/scrolless-backend/internal/user/repository"
	"github.com/dgrijalva/jwt-go"
)

type UserService interface {
	ListUser() (*userEntity.UserListResponse, error)
	GetUserByID(userId int) (*userEntity.UserDetailResponse, error)
	RegisterUser(user *userEntity.UserRegister) error
	LoginUser(user *userEntity.UserLogin) (string, error)
}

type userService struct {
	userRepo userRepo.UserRepository
}

func (service *userService) ListUser() (*userEntity.UserListResponse, error) {
	users, err := service.userRepo.ListUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *userService) GetUserByID(userId int) (*userEntity.UserDetailResponse, error) {
	user, err := service.userRepo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) RegisterUser(user *userEntity.UserRegister) error {
	return service.userRepo.RegisterUser(user)
}

func (service *userService) LoginUser(user *userEntity.UserLogin) (string, error) {
	dbUser, err := service.userRepo.LoginUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"apps":    "scrolless-backend",
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretToken := os.Getenv("JWT_SECRET")
	if secretToken == "" {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	stringToken, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return stringToken, nil

}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
