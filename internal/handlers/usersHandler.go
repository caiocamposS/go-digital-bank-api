package handlers

import (
	"digital-bank-api/internal/dto/request"
	"digital-bank-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	// get the email and password
	var req request.UserRequest

	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	userResponse, err := h.userService.CreateUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, userResponse)

}

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	resp, err := h.userService.Login(req)

	if err != nil {
		c. JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to login",
		})

		return
	}

	c.JSON(http.StatusOK, resp)

}
