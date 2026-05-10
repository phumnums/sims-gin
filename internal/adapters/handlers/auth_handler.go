package handlers

import (
	"cims/internal/adapters/dto"
	authPort "cims/internal/core/ports/auth"
	"cims/pkg/utils/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service authPort.Service
}

func NewAuthHandler(service authPort.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	token, status, err := h.service.Login(req)
	if err != nil {
		response.Error(c, status, err)
		return
	}

	response.Success(c, http.StatusOK, token, "login successfully")

}
