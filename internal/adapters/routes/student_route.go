package routes

import (
	studentHandler "cims/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterrStudentRoutes(r *gin.Engine, handler *studentHandler.StudentHandler) {
	students := r.Group("/student")
	{
		students.GET("", handler.GetStudents)
	}
}
