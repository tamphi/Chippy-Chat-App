package user

import (
	"context"
	"fmt"
	"server/internal/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	Repository
	timeout time.Duration
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// todo: secret key
const (
	secretKey = "secret"
)

func NewService(repository Repository) Service {
	return &service{Repository: repository, timeout: time.Duration(2) * time.Second}
}

func (s *service) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	returnedContext, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, _ := util.HashPassword(request.Password)
	u := &User{
		Username: request.Username,
		Password: hashedPassword,
	}

	newUser, err := s.Repository.CreateUser(returnedContext, u)

	if err != nil {
		return nil, fmt.Errorf("failed to create new user: %s", err)
	}

	response := &CreateUserResponse{
		ID:       strconv.Itoa(newUser.ID), //<-- typecaste int to str
		Username: newUser.Username,
		// Password: hashedPassword, // <-- uncomment for debugging
	}

	return response, nil

}

func (s *service) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	returnedContext, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	existingUser, err := s.Repository.GetUsername(returnedContext, request.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing username from repository: %s", err)
	}

	err = util.CheckPasswordHash(request.Password, existingUser.Password)

	if err != nil {
		return nil, fmt.Errorf("failed to find password hash: %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(existingUser.ID),
		Username: existingUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(existingUser.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %s", err)
	}

	return &LoginResponse{accessToken: ss, Username: existingUser.Username, ID: strconv.Itoa(existingUser.ID)}, nil
}

func (s *service) ListUsers(ctx context.Context) ([]*User, error) {
	returnedContext, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	users, err := s.Repository.ListUsers(returnedContext)

	if err != nil {
		return nil, fmt.Errorf("failed to list users: %s", err)
	}

	return users, nil

}
