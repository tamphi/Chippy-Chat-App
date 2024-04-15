package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service
}

func NewController(s Service) *Controller {
	return &Controller{Service: s}
}

func (controllerInstance *Controller) CreateUser(ginctx *gin.Context) {
	var userRequest CreateUserRequest
	err := ginctx.ShouldBindJSON(&userRequest)

	if err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := controllerInstance.Service.CreateUser(ginctx.Request.Context(), &userRequest)

	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginctx.JSON(http.StatusOK, userResponse)
}

func (controllerInstance *Controller) Login(ginctx *gin.Context) {
	var loginRequest LoginRequest
	err := ginctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := controllerInstance.Service.Login(ginctx.Request.Context(), &loginRequest)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginctx.SetCookie("jwt", loginResponse.accessToken, 3600, "/", "localhost", false, true)

	ginctx.JSON(http.StatusOK, gin.H{"login": loginResponse, "message": "login successful"})
}

func (controllerInstance *Controller) ListUsers(ginctx *gin.Context) {
	users, err := controllerInstance.Service.ListUsers(ginctx.Request.Context())
	var usernames []string
	if err != nil {
		return
	}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	ginctx.JSON(http.StatusOK, gin.H{"users": usernames, "message": "login successful"})
}

func (controllerInstance *Controller) Logout(ginctx *gin.Context) {
	ginctx.SetCookie("jwt", "", -1, "", "", false, true)
	ginctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
