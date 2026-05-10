package main

import (
	"cims/internal/adapters/handlers"
	"cims/internal/adapters/repositories"
	"cims/internal/adapters/routes"
	authService "cims/internal/core/services/auth"
	studentService "cims/internal/core/services/student"
	"cims/internal/infarstructure/config"
	"cims/internal/infarstructure/database"
	"log"
)

func main() {

	// config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init layer
	// student
	studentRepo := repositories.NewStudentRepository(db)
	studentService := studentService.NewStudentService(studentRepo)
	studentHandler := handlers.NewStudentHandler(studentService)
	// user
	userRepo := repositories.NewAuthRepository(db)
	userService := authService.NewAuthService(userRepo)
	userHandler := handlers.NewAuthHandler(userService)

	// setup routes
	r := routes.SetupRouter()
	routes.RegisterrStudentRoutes(r, studentHandler)
	routes.RegisterUserRoutes(r, userHandler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
