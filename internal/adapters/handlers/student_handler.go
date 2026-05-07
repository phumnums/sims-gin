package handlers

import (
	studentDTO "cims/internal/adapters/dto"
	studentPort "cims/internal/core/ports/student"
	"cims/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service studentPort.Service
}

func NewHandler(service studentPort.Service) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {

	var params studentDTO.SearchStudents
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	students, page, size, total, status, err := h.service.GetAll(params)
	if err != nil {
		response.Error(c, status, err.Error())
		return
	}
	response.SuccessPagination(c, students, status, page, size, total)
}
