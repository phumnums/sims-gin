package routes

import (
	"cims/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterStaffRoutes(r *gin.Engine, handler *handlers.StaffHandler) {
	students := r.Group("/staff")
	{
		students.POST("/", handler.Create)
	}
}
