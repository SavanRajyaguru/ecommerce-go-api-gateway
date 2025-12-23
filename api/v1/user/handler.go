package user

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	resp, err := h.service.Login(req)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Login successful", resp)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "User details", user)
}
