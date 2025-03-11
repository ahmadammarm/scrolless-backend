package entity

// Request
type UserRegister struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required, min=3, max=100"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, min=8"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}

type EditUser struct {
	Name     string `json:"name" validate:"required, min=3, max=100"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}

// Response
type UserJWT struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetailResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserListSubResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserListResponse struct {
	Users []UserListSubResponse `json:"users"`
}
