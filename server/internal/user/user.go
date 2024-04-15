package user

import "context"

type User struct {
	ID       int    `json:"id" db:"id"` // <--- use bigserial (64bit) for large user pool
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUsername(ctx context.Context, username string) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID       string `json:"id" db:"id"` // <--- don't forget to typecast
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type LoginRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	ID          string `json:"id" db:"id"` // <--- don't forget to typecast
	Username    string `json:"username" db:"username"`
	accessToken string
}

type Service interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error)
	Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
	ListUsers(ctx context.Context) ([]*User, error)
}
