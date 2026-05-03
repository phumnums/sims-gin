package handlers

import (
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
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	nationalID := c.Query("national_id")
	phoneNumber := c.Query("phone_number")
	email := c.Query("email")
	campusID := c.Query("campus_id")

	students, err := h.service.SearchStudents(
		firstName,
		lastName,
		nationalID,
		phoneNumber,
		email,
		campusID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}
