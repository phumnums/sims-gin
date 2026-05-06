package main

import (
	studentHandler "cims/internal/adapters/handlers"
	studentRepo "cims/internal/adapters/repositories"
	"cims/internal/adapters/routes"
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
	repo := studentRepo.NewStudentRepository(db)
	service := studentService.NewService(repo)
	handler := studentHandler.NewHandler(service)

	// setup routes
	r := routes.SetupRouter()
	routes.RegisterrStudentRoutes(r, handler)

	r.GET("/students", handler.SearchStudents)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
