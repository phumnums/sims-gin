package main

import (
	"cims/internal/adapters/handlers"
	"cims/internal/adapters/repositories"
	"cims/internal/adapters/routes"

	authService "cims/internal/core/services/auth"
	staffService "cims/internal/core/services/staff"
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
	// auth
	authRepo := repositories.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)
	// staff
	staffRepo := repositories.NewStaffRepository(db)
	staffService := staffService.NewStaffService(staffRepo, db)
	staffHandler := handlers.NewStaffHandler(staffService)

	// setup routes
	r := routes.SetupRouter()
	routes.RegisterrStudentRoutes(r, studentHandler)
	routes.RegisterAuthRoutes(r, authHandler)
	routes.RegisterStaffRoutes(r, staffHandler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
