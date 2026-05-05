package main

import (
	studentHandler "cims/internal/adapters/handlers"
	studentRepo "cims/internal/adapters/repositories"
	studentService "cims/internal/core/services/student"
	"cims/internal/infarstructure/config"
	"cims/internal/infarstructure/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := studentRepo.NewStudentRepository(db)
	service := studentService.NewService(repo)
	handler := studentHandler.NewHandler(service)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/students", handler.SearchStudents)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
