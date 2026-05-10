package routes

import (
	"cims/internal/adapters/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	fmt.Println("register user routes")
	r.POST("/login", handler.Login)
}
