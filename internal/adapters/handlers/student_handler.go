package handlers

import (
	studentDTO "cims/internal/adapters/dto"
	studentPort "cims/internal/core/ports/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service studentPort.Service
}

func NewHandler(service studentPort.Service) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) SearchStudents(c *gin.Context) {

	var params studentDTO.SearchStudents
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	students, page, size, total, err := h.service.SearchStudents(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"page":  page,
		"size":  size,
		"total": total,
	})
}
