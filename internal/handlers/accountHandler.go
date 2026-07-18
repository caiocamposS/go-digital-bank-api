package handlers

import (
	"digital-bank-api/internal/dto/request"
	"digital-bank-api/internal/dto/response"
	"digital-bank-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: service,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req request.AccountRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})

		return
	}

	var resp *response.AccountResponse

	resp, err = h.accountService.CreateAccount(&req, userID.(uint))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create account",
		})

		return
	}

	c.JSON(http.StatusCreated, resp)
}
