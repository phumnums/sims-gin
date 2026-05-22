package handlers

import (
	"cims/internal/adapters/dto"
	"cims/internal/adapters/middleware"
	staffPort "cims/internal/core/ports/staff"
	"cims/pkg/utils/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	service staffPort.Service
}

func NewStaffHandler(service staffPort.Service) *StaffHandler {
	return &StaffHandler{service: service}
}

func (h StaffHandler) Create(c *gin.Context) {
	var req dto.CreateStaffRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	ctx := c.Request.Context()
	val, _ := c.Get("auth")
	authCtx := val.(middleware.AuthContext)

	status, err := h.service.Create(ctx, authCtx, req)
	if err != nil {
		response.Error(c, status, err)
		return
	}

}
