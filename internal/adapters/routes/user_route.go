package routes

import (
	"cims/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	r.POST("/login", handler.Login)
}
