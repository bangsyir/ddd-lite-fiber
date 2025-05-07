package main

import (
	"DDD-fiberv2/internal/domain/user"
	"DDD-fiberv2/internal/infrastructure/db"
	repository "DDD-fiberv2/internal/repository/user"
	http "DDD-fiberv2/internal/transport/http/user"
	"log"

	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initialize sqlite database
	sqlDB, err := db.NewSQLiteDB("./test.sqlite")

	if err != nil {
		log.Fatal("Failed to connect to SQLIite: %w", err)
	}
	defer sqlDB.Close()

	app := fiber.New()

	// initialize dependencies
	repo := repository.NewSQLiteRepository(sqlDB)
	svc := user.NewService(repo)
	handler := http.NewHandler(svc)

	// Register routes
	handler.RegisterRoute(app)
	// Initialize default config (Assign the middleware to /metrics)
	app.Get("/metrics", monitor.New())
	// Start server

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server : %w", err)
	}
}
