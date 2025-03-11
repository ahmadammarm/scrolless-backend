package handler

import (
    userService "github.com/ahmadammarm/scrolless-backend/internal/user/service"
)

type userHandler struct {
    userService userService.UserService
    
}